package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"ordering-platform/api/admin/internal/svc"
	contextUtil "ordering-platform/pkg/utils/context"
	"ordering-platform/rpc/sys/sysclient"
)

type UserCodesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCodesLogic {
	return &UserCodesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCodesLogic) UserCodes() (resp []string, err error) {
	// todo: add your logic here and delete this line
	// 获取当前用户 所有按钮级别权限
	roleId, err := contextUtil.GetRoleIDFromContext(l.ctx)
	if err != nil {
		return nil, err
	}

	menus, err := l.svcCtx.MenuService.ListMenuByRole(l.ctx, &sysclient.ListMenuRoleReq{
		RoleId: roleId,
	})

	if err != nil {
		return nil, err
	}
	for _, menu := range menus.Data {
		if menu.MenuType == "F" {
			resp = append(resp, menu.Permission)
		}
	}

	return resp, nil
}
