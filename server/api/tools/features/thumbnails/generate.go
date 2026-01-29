package thumbnails

import (
	"bytes"
	"context"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
)

const (
	ThumbnailMaxSize = 400
	ThumbnailSuffix  = "_thumb"
	ThumbnailQuality = 80
)

type Generator struct {
	minio           *minio.Client
	bucket          string
	presignedExpiry time.Duration
}

func NewGenerator(minioClient *minio.Client, bucket string, presignedExpiry time.Duration) *Generator {
	return &Generator{
		minio:           minioClient,
		bucket:          bucket,
		presignedExpiry: presignedExpiry,
	}
}

// GetThumbnailKey returns the S3 key for a thumbnail given the original key
func GetThumbnailKey(originalKey string) string {
	return originalKey + ThumbnailSuffix
}

// ThumbnailExists checks if a thumbnail exists for the given original key
func (g *Generator) ThumbnailExists(ctx context.Context, originalKey string) bool {
	thumbKey := GetThumbnailKey(originalKey)
	_, err := g.minio.StatObject(ctx, g.bucket, thumbKey, minio.StatObjectOptions{})
	return err == nil
}

// GetThumbnailURL returns a presigned URL for the thumbnail if it exists
func (g *Generator) GetThumbnailURL(ctx context.Context, originalKey string) (*url.URL, error) {
	thumbKey := GetThumbnailKey(originalKey)

	// Check if thumbnail exists
	if !g.ThumbnailExists(ctx, originalKey) {
		return nil, nil
	}

	return g.minio.PresignedGetObject(ctx, g.bucket, thumbKey, g.presignedExpiry, nil)
}

// GenerateThumbnail downloads the original image, creates a thumbnail, and uploads it
func (g *Generator) GenerateThumbnail(ctx context.Context, originalKey string, mimeType string) error {
	// Only process images
	if !strings.HasPrefix(mimeType, "image/") {
		return nil
	}

	// Skip if thumbnail already exists
	if g.ThumbnailExists(ctx, originalKey) {
		return nil
	}

	// Download original
	obj, err := g.minio.GetObject(ctx, g.bucket, originalKey, minio.GetObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to get original: %w", err)
	}
	defer obj.Close()

	// Read all data into memory so we can decode with EXIF orientation
	data, err := io.ReadAll(obj)
	if err != nil {
		return fmt.Errorf("failed to read image data: %w", err)
	}

	// Decode image with automatic EXIF orientation correction
	img, err := imaging.Decode(bytes.NewReader(data), imaging.AutoOrientation(true))
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// Resize to fit within ThumbnailMaxSize while maintaining aspect ratio
	thumbnail := imaging.Fit(img, ThumbnailMaxSize, ThumbnailMaxSize, imaging.Lanczos)

	// Encode thumbnail
	var buf bytes.Buffer
	var contentType string

	// Determine format from mime type
	if strings.Contains(mimeType, "png") {
		err = png.Encode(&buf, thumbnail)
		contentType = "image/png"
	} else {
		// Use JPEG for other formats (including jpeg, gif, etc.)
		err = jpeg.Encode(&buf, thumbnail, &jpeg.Options{Quality: ThumbnailQuality})
		contentType = "image/jpeg"
	}

	if err != nil {
		return fmt.Errorf("failed to encode thumbnail: %w", err)
	}

	// Upload thumbnail
	thumbKey := GetThumbnailKey(originalKey)
	_, err = g.minio.PutObject(ctx, g.bucket, thumbKey, &buf, int64(buf.Len()), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return fmt.Errorf("failed to upload thumbnail: %w", err)
	}

	return nil
}

// EnsureThumbnail generates a thumbnail if it doesn't exist and returns its URL
func (g *Generator) EnsureThumbnail(ctx context.Context, originalKey string, mimeType string) (*url.URL, error) {
	// Only process images
	if !strings.HasPrefix(mimeType, "image/") {
		return nil, nil
	}

	// Generate if doesn't exist
	if err := g.GenerateThumbnail(ctx, originalKey, mimeType); err != nil {
		return nil, err
	}

	return g.GetThumbnailURL(ctx, originalKey)
}

// GenerateThumbnailAsync generates a thumbnail in the background
func (g *Generator) GenerateThumbnailAsync(originalKey string, mimeType string) {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_ = g.GenerateThumbnail(ctx, originalKey, mimeType)
	}()
}
