package role

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// role详情
func NewRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleInfoLogic {
	return &RoleInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleInfoLogic) RoleInfo(req *types.RoleInfoReq) (resp *types.RoleInfoData, err error) {
	// todo: add your logic here and delete this line
	info, err := l.svcCtx.RoleService.RoleInfo(l.ctx, &sysclient.RoleInfoReq{
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, err
	}

	// 构建menuTree

	return &types.RoleInfoData{
		RoleId:        info.RoleId,
		RoleName:      info.RoleName,
		RoleKey:       info.RoleKey,
		Status:        info.Status,
		Sort:          info.Sort,
		DefaultRouter: info.DefaultRouter,
		SelectMenus:   info.SelectMenus,
		Admin:         info.Admin,
	}, nil
}
