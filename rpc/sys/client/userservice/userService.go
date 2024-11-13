// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: sys.proto

package userservice

import (
	"context"

	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginReq  = sysclient.LoginReq
	LoginResp = sysclient.LoginResp

	UserService interface {
		// 用户登录
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

// 用户登录
func (m *defaultUserService) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}
