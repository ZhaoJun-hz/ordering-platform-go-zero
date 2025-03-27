package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"ordering-platform/api/admin/internal/config"
	"ordering-platform/api/admin/internal/middleware"
	customcasbin "ordering-platform/pkg/casbin"
	"ordering-platform/rpc/sys/client/apiservice"
	"ordering-platform/rpc/sys/client/deptservice"
	"ordering-platform/rpc/sys/client/menuservice"
	"ordering-platform/rpc/sys/client/roleservice"
	"ordering-platform/rpc/sys/client/tokenservice"
	"ordering-platform/rpc/sys/client/userservice"
)

type ServiceContext struct {
	Config           config.Config
	AuthCheckRole    rest.Middleware
	PermissionAction rest.Middleware
	Redis            *redis.Redis
	Casbin           *casbin.Enforcer
	UserService      userservice.UserService
	ApiService       apiservice.ApiService
	MenuService      menuservice.MenuService
	DeptService      deptservice.DeptService
	RoleService      roleservice.RoleService
	TokenService     tokenservice.TokenService
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(c.Redis)
	sysClient := zrpc.MustNewClient(c.SysRpc)
	//DB, err := gorm.Open(mysql.Open(c.Mysql.Datasource), &gorm.Config{
	//	SkipDefaultTransaction: true,
	//	PrepareStmt:            true,
	//	Logger:                 settingLogConfig(),
	//})
	//if err != nil {
	//	panic(err)
	//}

	enforcer := customcasbin.Init(c.Mysql.Datasource, c.Redis)

	logx.Debug("mysql已连接")
	return &ServiceContext{
		Config:           c,
		AuthCheckRole:    middleware.NewAuthCheckRoleMiddleware().Handle,
		PermissionAction: middleware.NewPermissionActionMiddleware().Handle,

		Casbin: enforcer,
		Redis:  rds,

		UserService:  userservice.NewUserService(sysClient),
		ApiService:   apiservice.NewApiService(sysClient),
		MenuService:  menuservice.NewMenuService(sysClient),
		DeptService:  deptservice.NewDeptService(sysClient),
		RoleService:  roleservice.NewRoleService(sysClient),
		TokenService: tokenservice.NewTokenService(sysClient),
	}
}
