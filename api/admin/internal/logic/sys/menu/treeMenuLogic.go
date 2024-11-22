package menu

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TreeMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTreeMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TreeMenuLogic {
	return &TreeMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TreeMenuLogic) TreeMenu(req *types.TreeMenuReq) (resp *types.TreeMenuResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.MenuService.ListMenu(l.ctx, &sysclient.ListMenuReq{})

	if err != nil {
		return nil, err
	}
	var treeData = make(map[int64]*types.TreeMenuData)
	for _, data := range result.Data {
		if data.MenuType == "F" && !req.NeedButton {
			continue
		}
		treeData[data.MenuId] = &types.TreeMenuData{
			MenuId:       data.MenuId,
			MenuType:     data.MenuType,
			ParentMenuId: data.ParentMenuId,
			Name:         data.Name,
			Component:    data.Component,
			Path:         data.Path,
			Permission:   data.Permission,
			LinkFlag:     data.LinkFlag,
			Title:        data.Title,
			Meta: &types.TreeMenuMeta{
				Title:                    data.Title,
				Icon:                     data.Icon,
				KeepAlive:                data.IgnoreKeepAlive,
				HideInMenu:               data.HideInMenu,
				Link:                     data.Link,
				MenuVisibleWithForbidden: data.Disabled,
				Order:                    data.Sort,
			},
			Children: nil,
		}
	}

	var tree []*types.TreeMenuData

	// Create a map to efficiently find and attach children.
	for _, node := range treeData {
		if node.ParentMenuId == 0 {
			// Root node, add to tree.
			tree = append(tree, node)
		} else {
			// Find parent node and attach this node as a child.
			if parent, found := treeData[node.ParentMenuId]; found {
				parent.Children = append(parent.Children, node)
			}
		}
	}

	return &types.TreeMenuResp{
		Data: tree,
	}, nil
}
