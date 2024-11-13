// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package handler

import (
	"net/http"

	sysuser "ordering-platform/api/admin/internal/handler/sys/user"
	"ordering-platform/api/admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthCheckRole, serverCtx.PermissionAction},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/codes",
					Handler: sysuser.UserCodesHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/info",
					Handler: sysuser.UserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/menus",
					Handler: sysuser.UserMenusHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/sys/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/user/login",
				Handler: sysuser.UserLoginHandler(serverCtx),
			},
		},
	)
}