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
	"strings"

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
	printTree(tree, 0)

	var selectMenuIds = make(map[int64]bool, 0)
	for _, roleMenu := range roleMenus {
		selectMenuIds[roleMenu.MenuID] = true
	}

	filteredMenus := filterMenus(tree, selectMenuIds)

	realSelectMenuIds := getAllMenuIds(filteredMenus)

	return &sysclient.RoleInfoResp{
		RoleId:        role.RoleID,
		RoleName:      role.RoleName,
		RoleKey:       role.RoleKey,
		Status:        role.Status,
		Sort:          role.Sort,
		SelectMenus:   realSelectMenuIds,
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
func filterMenus(menuTree []*MenuTree, selectedIDs map[int64]bool) []*MenuTree {
	var filtered []*MenuTree

	for _, menu := range menuTree {
		if len(menu.Children) == 0 {
			// 叶子节点：直接检查是否被选中
			if selectedIDs[menu.MenuId] {
				filtered = append(filtered, menu)
			}
		} else {
			// 非叶子节点：递归过滤子菜单
			menu.Children = filterMenus(menu.Children, selectedIDs)

			// 检查所有子菜单是否都被选中
			allChildrenSelected := true
			for _, child := range menu.Children {
				if !selectedIDs[child.MenuId] {
					allChildrenSelected = false
					break
				}
			}

			// 如果当前菜单被选中，且所有子菜单都被选中，则保留父菜单
			if allChildrenSelected && selectedIDs[menu.MenuId] {
				filtered = append(filtered, menu)
			} else if len(menu.Children) > 0 {
				// 如果当前菜单未满足条件，但部分子菜单满足，保留子菜单
				filtered = append(filtered, menu.Children...)
			}
		}
	}
	return filtered
}

func getAllMenuIds(menus []*MenuTree) []int64 {
	var result []int64

	// 定义递归函数
	var collectMenuIds func(menu *MenuTree)
	collectMenuIds = func(menu *MenuTree) {
		// 添加当前菜单的 MenuId
		result = append(result, menu.MenuId)

		// 递归处理子菜单
		for _, child := range menu.Children {
			collectMenuIds(child)
		}
	}

	// 遍历菜单列表
	for _, menu := range menus {
		collectMenuIds(menu)
	}
	return result
}

func printTree(tree []*MenuTree, level int) {
	for _, menu := range tree {
		fmt.Printf("%sMenuId: %d, ParentMenuId: %d\n", strings.Repeat("  ", level), menu.MenuId, menu.ParentMenuId)
		if len(menu.Children) > 0 {
			printTree(menu.Children, level+1)
		}
	}
}
