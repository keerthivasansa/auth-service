package auth_grpc

import (
	context "context"
	"kv/auth/auth"

	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) IsAuth(ctx context.Context, req *AuthRequest) (*AuthResponse, error) {
	result := AuthResponse{
		Authenticated: req.Token == "",
	}
	return &result, nil
}

func (s *Server) GenerateToken(ctx context.Context, req *AuthTokenRequest) (*AuthToken, error) {
	userId, err := auth.SignInUser(ctx, req.Email, req.Password)
	token := &AuthToken{
		Token: userId,
	}
	return token, err
}

func RegisterAuthGrpc(server *grpc.Server) {
	authServer := &Server{}
	RegisterAuthServiceServer(server, authServer)
}
