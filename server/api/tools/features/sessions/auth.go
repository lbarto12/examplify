package sessions

import (
	"context"
	"server/api/tools/internaltools/passwords"
	"server/api/tools/internaltools/webtokens"
	"server/api/validation"
	"server/sqlc/sqlgen"
)

// SignIn authenticates a user with email and password, returning a JWT token
func (mgr SessionManager) SignIn(ctx context.Context, email string, password string) (*SessionResult, error) {
	// Validate email format
	if err := validation.ValidateEmail(email); err != nil {
		return nil, err
	}

	// Get user by email
	user, err := mgr.Queries.GetUserAccountByEmail(ctx, email)
	if err != nil {
		return nil, ErrUserNotFound
	}

	// Verify password
	if err := passwords.CompareHashAndPassword(password, user.PasswordHash); err != nil {
		return nil, ErrInvalidPassword
	}

	// Generate JWT token
	token, _, err := webtokens.GenerateJWT(user.ID)
	if err != nil {
		return nil, ErrTokenGeneration
	}

	return &SessionResult{
		Token:  token,
		UserID: user.ID,
	}, nil
}

// SignUp creates a new user account and returns a JWT token
func (mgr SessionManager) SignUp(ctx context.Context, email string, password string) (*SessionResult, error) {
	// Validate email format
	if err := validation.ValidateEmail(email); err != nil {
		return nil, err
	}

	// Validate password is non-empty
	if err := validation.ValidateNonEmpty("password", password); err != nil {
		return nil, err
	}

	// Hash password
	hashedPassword, err := passwords.GenerateFromPassword(password, passwords.NewDefaultPasswordGenerationOptions())
	if err != nil {
		return nil, err
	}

	// Create account
	user, err := mgr.Queries.CreateAccount(ctx, sqlgen.CreateAccountParams{
		UserEmail:    email,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		return nil, ErrAccountCreation
	}

	// Generate JWT token
	token, _, err := webtokens.GenerateJWT(user.ID)
	if err != nil {
		return nil, ErrTokenGeneration
	}

	return &SessionResult{
		Token:  token,
		UserID: user.ID,
	}, nil
}
