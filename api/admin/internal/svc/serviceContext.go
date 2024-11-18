package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"ordering-platform/api/admin/internal/config"
	"ordering-platform/api/admin/internal/middleware"
	"ordering-platform/rpc/sys/client/apiservice"
	"ordering-platform/rpc/sys/client/menuservice"
	"ordering-platform/rpc/sys/client/userservice"
)

type ServiceContext struct {
	Config           config.Config
	AuthCheckRole    rest.Middleware
	PermissionAction rest.Middleware

	UserService userservice.UserService
	ApiService  apiservice.ApiService
	MenuService menuservice.MenuService
}

func NewServiceContext(c config.Config) *ServiceContext {

	sysClient := zrpc.MustNewClient(c.SysRpc)
	return &ServiceContext{
		Config:           c,
		AuthCheckRole:    middleware.NewAuthCheckRoleMiddleware().Handle,
		PermissionAction: middleware.NewPermissionActionMiddleware().Handle,

		UserService: userservice.NewUserService(sysClient),
		ApiService:  apiservice.NewApiService(sysClient),
		MenuService: menuservice.NewMenuService(sysClient),
	}
}
