package validation

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

// ValidateUUID parses and validates a UUID string
func ValidateUUID(id string) (uuid.UUID, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, errors.New("invalid UUID format")
	}
	return parsed, nil
}

// ValidateNonEmpty checks if a string field is non-empty
func ValidateNonEmpty(field, value string) error {
	if strings.TrimSpace(value) == "" {
		return errors.New(field + " cannot be empty")
	}
	return nil
}

// ValidateMimeType validates that the MIME type is one of the allowed types
func ValidateMimeType(mimeType string) error {
	allowed := []string{"application/pdf", "image/png", "image/jpeg", "image/jpg"}
	for _, allowedType := range allowed {
		if mimeType == allowedType {
			return nil
		}
	}
	return errors.New("unsupported mime type: " + mimeType)
}

// ValidateEmail performs basic email validation
func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return errors.New("email cannot be empty")
	}
	if !strings.Contains(email, "@") {
		return errors.New("invalid email format")
	}
	return nil
}
