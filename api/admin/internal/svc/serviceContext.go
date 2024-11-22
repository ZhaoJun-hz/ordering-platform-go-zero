package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"ordering-platform/api/admin/internal/config"
	"ordering-platform/api/admin/internal/middleware"
	customcasbin "ordering-platform/pkg/casbin"
	"ordering-platform/rpc/sys/client/apiservice"
	"ordering-platform/rpc/sys/client/deptservice"
	"ordering-platform/rpc/sys/client/menuservice"
	"ordering-platform/rpc/sys/client/roleservice"
	"ordering-platform/rpc/sys/client/userservice"
	"time"
)

type ServiceContext struct {
	Config           config.Config
	AuthCheckRole    rest.Middleware
	PermissionAction rest.Middleware

	Casbin *casbin.SyncedEnforcer

	UserService userservice.UserService
	ApiService  apiservice.ApiService
	MenuService menuservice.MenuService
	DeptService deptservice.DeptService
	RoleService roleservice.RoleService
}

func NewServiceContext(c config.Config) *ServiceContext {

	sysClient := zrpc.MustNewClient(c.SysRpc)
	DB, err := gorm.Open(mysql.Open(c.Mysql.Datasource), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 settingLogConfig(),
	})
	if err != nil {
		panic(err)
	}

	enforcer := customcasbin.Init(DB)

	logx.Debug("mysql已连接")
	return &ServiceContext{
		Config:           c,
		AuthCheckRole:    middleware.NewAuthCheckRoleMiddleware().Handle,
		PermissionAction: middleware.NewPermissionActionMiddleware().Handle,

		Casbin: enforcer,

		UserService: userservice.NewUserService(sysClient),
		ApiService:  apiservice.NewApiService(sysClient),
		MenuService: menuservice.NewMenuService(sysClient),
		DeptService: deptservice.NewDeptService(sysClient),
		RoleService: roleservice.NewRoleService(sysClient),
	}
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	logx.Infof(format, args...)
}

// init log config
func settingLogConfig() logger.Interface {
	newLogger := logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)
	return newLogger
}
