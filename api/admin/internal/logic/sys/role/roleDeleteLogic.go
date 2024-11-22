package role

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// role删除
func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDeleteLogic) RoleDelete(req *types.RoleDeleteReq) (resp *types.RoleDeleteResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.RoleService.RoleDelete(l.ctx, &sysclient.RoleDeleteReq{
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.Casbin.RemoveFilteredPolicy(0, result.RoleKey)
	if err != nil {
		return nil, err
	}
	return
}
