package auth

import (
	"context"
	"errors"
	"kv/auth/db"
	"kv/auth/utils"
)

var ErrWrongCredentials = errors.New("invalid credentials")

func SignInUser(ctx context.Context, email string, password string) (string, error) {
	queries := db.GetQuerier()
	user, err := queries.GetUserWithEmail(ctx, email)
	if err != nil {
		return "", err
	}
	pwdHash := utils.Hash(password)
	if user.Password != pwdHash {
		return "", ErrWrongCredentials
	}
	return user.ID, nil
}
