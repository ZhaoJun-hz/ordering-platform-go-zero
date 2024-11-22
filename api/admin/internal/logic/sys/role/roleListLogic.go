package role

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// role列表
func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.RoleListReq) (resp *types.RoleListResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.RoleService.RoleList(l.ctx, &sysclient.RoleListReq{
		RoleName: req.RoleName,
		RoleKey:  req.RoleKey,
		Status:   req.Status,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var list []*types.RoleListData
	for _, item := range result.Data {
		list = append(list, &types.RoleListData{
			RoleId:        item.RoleId,
			RoleName:      item.RoleName,
			RoleKey:       item.RoleKey,
			Status:        item.Status,
			Sort:          item.Sort,
			DefaultRouter: item.DefaultRouter,
			CreateTime:    item.CreateTime,
		})
	}
	return &types.RoleListResp{
		Total: result.Total,
		Data:  list,
	}, nil
}
