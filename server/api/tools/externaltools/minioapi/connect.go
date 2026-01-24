package minioapi

import (
	"server/environment"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Connect(env *environment.Vars) (*minio.Client, error) {
	return minio.New(env.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(env.MinioUser, env.MinioPassword, ""),
		Secure: env.MinioUseSSL,
	})
}
