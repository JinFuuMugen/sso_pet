package auth

import (
	"context"
	sso_v1 "sso_pet/proto/go"

	"google.golang.org/grpc"
)

type serverAPI struct {
	sso_v1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	sso_v1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *sso_v1.LoginRequest) (*sso_v1.LoginResponse, error) {
	return &sso_v1.LoginResponse{
		Token: "token1234", //TODO: change
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *sso_v1.RegisterRequest) (*sso_v1.RegisterResponse, error) {
	panic("implement me") //TODO: change
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *sso_v1.IsAdminRequest) (*sso_v1.IsAdminResponse, error) {
	panic("implement me") //TODO: change
}
