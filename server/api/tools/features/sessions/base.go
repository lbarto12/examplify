package sessions

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"server/sqlc/sqlgen"
)

type session_manager_interface interface {
	SignIn(ctx context.Context, email string, password string) (*SessionResult, error)
	SignUp(ctx context.Context, email string, password string) (*SessionResult, error)
}

type SessionManager struct {
	Queries *sqlgen.Queries
}

type NewSessionManagerParams struct {
	Queries *sqlgen.Queries
}

func NewSessionManager(data NewSessionManagerParams) (*SessionManager, error) {
	if data.Queries == nil {
		return nil, ErrNoQueries
	}

	// Validate against interface
	var mgr session_manager_interface = &SessionManager{
		Queries: data.Queries,
	}
	return mgr.(*SessionManager), nil
}

type SessionResult struct {
	Token  string
	UserID uuid.UUID
}

var ( // Errors
	ErrNoQueries         error = errors.New("no queries provided")
	ErrInvalidEmail      error = errors.New("invalid email format")
	ErrEmptyPassword     error = errors.New("password cannot be empty")
	ErrUserNotFound      error = errors.New("user not found")
	ErrInvalidPassword   error = errors.New("invalid password")
	ErrAccountCreation   error = errors.New("failed to create account")
	ErrTokenGeneration   error = errors.New("failed to generate token")
)
