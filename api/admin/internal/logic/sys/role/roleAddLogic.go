package role

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// role添加
func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleAddLogic) RoleAdd(req *types.RoleAddReq) (resp *types.RoleAddResp, err error) {
	// todo: add your logic here and delete this line

	result, err := l.svcCtx.RoleService.RoleAdd(l.ctx, &sysclient.RoleAddReq{
		RoleName:      req.RoleName,
		RoleKey:       req.RoleKey,
		Status:        req.Status,
		Sort:          req.Sort,
		SelectMenus:   req.SelectMenus,
		DefaultRouter: req.DefaultRouter,
	})
	if err != nil {
		return nil, err
	}

	mp := make(map[string]interface{}, 0)
	polices := make([][]string, 0)
	for _, api := range result.Data {
		if mp[req.RoleKey+"-"+api.Path+"-"+api.Action] != "" {
			mp[req.RoleKey+"-"+api.Path+"-"+api.Action] = ""
			polices = append(polices, []string{req.RoleKey, api.Path, api.Action})
		}
	}

	if len(polices) > 0 {
		// 添加 casbin
		_, err := l.svcCtx.Casbin.AddNamedPolicies("p", polices)
		if err != nil {
			return nil, err
		}
	}
	err = l.svcCtx.Casbin.LoadPolicy()
	if err != nil {
		return nil, err
	}

	return
}
