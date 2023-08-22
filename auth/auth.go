package auth

import (
	"context"
	"errors"
	"kv/auth/db"
	"kv/auth/utils"
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
