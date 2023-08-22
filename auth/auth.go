package auth

import (
	"context"
	"errors"
	"kv/auth/db"
	"kv/auth/utils"

	"github.com/google/uuid"
)

var ErrWrongCredentials = errors.New("invalid credentials")

func SignInUser(ctx context.Context, provider string, providerId string, password string) (string, error) {
	queries := db.GetQuerier()
	user, err := queries.GetUser(ctx, db.GetUserParams{
		Provider:   provider,
		Providerid: providerId,
	})
	if err != nil {
		return "", err
	}
	pwdHash := utils.Hash(password)
	if user.Password != pwdHash {
		return "", ErrWrongCredentials
	}
	return user.ID, nil
}

func CreateUser(ctx context.Context, provider string, providerId string, password string) error {
	queries := db.GetQuerier()
	userId := uuid.NewString()
	return queries.CreateUser(ctx, db.CreateUserParams{
		ID:         userId,
		Password:   utils.Hash(password),
		Provider:   provider,
		Providerid: providerId,
	})
}
