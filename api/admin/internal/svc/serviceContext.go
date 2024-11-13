package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"ordering-platform/api/admin/internal/config"
	"ordering-platform/api/admin/internal/middleware"
	"ordering-platform/rpc/sys/client/userservice"
)

type ServiceContext struct {
	Config           config.Config
	AuthCheckRole    rest.Middleware
	PermissionAction rest.Middleware

	UserService userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {

	sysClient := zrpc.MustNewClient(c.SysRpc)
	return &ServiceContext{
		Config:           c,
		AuthCheckRole:    middleware.NewAuthCheckRoleMiddleware().Handle,
		PermissionAction: middleware.NewPermissionActionMiddleware().Handle,

		UserService: userservice.NewUserService(sysClient),
	}
}
