package auth

import (
	context "context"

	"google.golang.org/grpc"
)

type Server struct {
}

func (*Server) IsAuth(ctx context.Context, req *AuthRequest) (*AuthResponse, error) {
	result := AuthResponse{
		Authenticated: req.Token == "",
	}
	return &result, nil
}

func (*Server) GenerateToken(ctx context.Context, req *AuthTokenRequest) (*AuthToken, error) {
	provider := req.GetProvider()
	if provider != AuthProvider_EMAIL {
		return nil, ErrUnsupportedProvider
	}
	userId, err := SignInUser(ctx, provider.String(), req.ProviderId, req.Password)
	token := &AuthToken{
		Token: userId,
	}
	return token, err
}

func (*Server) SignUp(ctx context.Context, req *AuthTokenRequest) (*Result, error) {
	res := Result{
		Okay: true,
	}
	err := CreateUser(ctx, req.Provider.String(), req.ProviderId, req.Password)
	return &res, err
}

func RegisterAuthGrpc(server *grpc.Server) {
	authServer := &Server{}
	RegisterAuthServiceServer(server, authServer)
}
