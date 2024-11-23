package roleservicelogic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/errcode"
	"ordering-platform/rpc/sys/gen/query"
	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleInfoLogic {
	return &RoleInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleInfoLogic) RoleInfo(in *sysclient.RoleInfoReq) (*sysclient.RoleInfoResp, error) {
	// todo: add your logic here and delete this line
	role, err := query.SysRole.WithContext(l.ctx).Where(query.SysRole.RoleID.Eq(in.RoleId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "Role不存在,参数：%d,异常:%s", in.RoleId, err.Error())
		return nil, errors.WithStack(errcode.RoleNotExistError)
	}
	if err != nil {
		logc.Errorf(l.ctx, "查询 Role失败,参数：%d,异常:%s", in.RoleId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Menu失败 %v, param %v", err, in.RoleId)
	}

	// 角色关联的菜单
	roleMenus, err := query.SysRoleMenu.WithContext(l.ctx).Where(query.SysRoleMenu.RoleID.Eq(in.RoleId)).Find()
	if err != nil {
		logc.Errorf(l.ctx, "查询 SysRoleMenu 列表失败,参数：%d,异常:%s", in.RoleId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 SysRoleMenu 列表失败 %v, param %v", err, in.RoleId)
	}

	// 获取所有的菜单
	menus, err := query.SysMenu.WithContext(l.ctx).Order(query.SysMenu.Sort).Find()
	if err != nil {
		logc.Errorf(l.ctx, "查询 Menu 列表失败, 异常:%s", err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Menu 列表失败 %v", err)
	}
	// 构建树
	var treeData []MenuTree
	for _, menu := range menus {
		fmt.Printf("MenuId: %d, ParentMenuId: %d\n", menu.MenuID, menu.ParentMenuID)
		treeData = append(treeData, MenuTree{
			MenuId:       menu.MenuID,
			ParentMenuId: menu.ParentMenuID,
		})
	}
	tree := buildMenuTree(treeData, 0)

	var selectMenuIds = make(map[int64]bool, 0)
	for _, roleMenu := range roleMenus {
		selectMenuIds[roleMenu.MenuID] = true
	}

	filteredMenus := filterMenus(tree, selectMenuIds)
	// 仅包含 31 41
	//realSelectMenuIds := getAllMenuIds(filteredMenus)

	return &sysclient.RoleInfoResp{
		RoleId:        role.RoleID,
		RoleName:      role.RoleName,
		RoleKey:       role.RoleKey,
		Status:        role.Status,
		Sort:          role.Sort,
		SelectMenus:   filteredMenus,
		DefaultRouter: role.DefaultRouter,
		Admin:         role.Admin,
	}, nil
}

type MenuTree struct {
	MenuId       int64       // 菜单ID
	ParentMenuId int64       // 父菜单ID
	Children     []*MenuTree // 子菜单
}

// 构建菜单树
func buildMenuTree(menus []MenuTree, parentMenuId int64) []*MenuTree {
	// 用于存储当前节点的子菜单
	var tree []*MenuTree

	// 遍历所有菜单
	for i := range menus {

		menu := menus[i] // 显式创建副本，避免覆盖问题

		// 如果当前菜单是 parentMenuId 的子菜单
		if menu.ParentMenuId == parentMenuId {
			fmt.Printf("Building tree: ParentMenuId: %d -> MenuId: %d\n", parentMenuId, menu.MenuId)
			// 递归查找子菜单
			menu.Children = buildMenuTree(menus, menu.MenuId)
			// 添加到树中
			tree = append(tree, &menu)
		}
	}

	return tree
}

// 过滤菜单，剔除未选中所有子菜单的父菜单
func filterMenus(menuTree []*MenuTree, selectedIDs map[int64]bool) []int64 {
	// 用于存储最终过滤后的选中菜单ID
	var result []int64

	// 辅助函数，对菜单树进行递归过滤
	var dfs func(node *MenuTree) bool
	dfs = func(node *MenuTree) bool {
		// 是否保留当前节点
		keep := selectedIDs[node.MenuId]

		// 如果当前节点有子节点，递归检查子节点
		if len(node.Children) > 0 {
			allChildrenSelected := true

			for _, child := range node.Children {
				childKeep := dfs(child) // 递归检查子节点是否保留
				if !childKeep {
					allChildrenSelected = false
				}
			}

			// 如果所有子节点都被选择，且当前节点被选择，则保留当前节点
			if !allChildrenSelected {
				keep = false // 如果有未选中的子节点，则当前节点不能保留
			}
		}

		// 如果当前节点被保留，则添加到结果中
		if keep {
			result = append(result, node.MenuId)
		}
		return keep
	}

	// 遍历根节点（即 menuTree 的每个顶级节点）
	for _, root := range menuTree {
		dfs(root)
	}

	return result
}
