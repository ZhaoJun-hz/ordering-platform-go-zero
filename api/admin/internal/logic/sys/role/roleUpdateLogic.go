package role

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// role更新
func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleUpdateLogic) RoleUpdate(req *types.RoleUpdateReq) (resp *types.RoleUpdateResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.RoleService.RoleUpdate(l.ctx, &sysclient.RoleUpdateReq{
		RoleId:        req.RoleId,
		RoleName:      req.RoleName,
		Status:        req.Status,
		Sort:          req.Sort,
		SelectMenus:   req.SelectMenus,
		DefaultRouter: req.DefaultRouter,
	})
	if err != nil {
		return nil, err
	}
	l.svcCtx.Casbin.RemoveFilteredPolicy(0, result.RoleKey)
	mp := make(map[string]interface{}, 0)
	polices := make([][]string, 0)
	for _, api := range result.Data {
		if mp[result.RoleKey+"-"+api.Path+"-"+api.Action] != "" {
			mp[result.RoleKey+"-"+api.Path+"-"+api.Action] = ""
			polices = append(polices, []string{result.RoleKey, api.Path, api.Action})
		}
	}

	if len(polices) > 0 {
		// 添加 casbin
		l.svcCtx.Casbin.AddNamedPolicies("p", polices)
	}
	l.svcCtx.Casbin.LoadPolicy()
	return
}
