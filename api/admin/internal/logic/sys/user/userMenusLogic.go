package user

import (
	"context"
	contextUtil "ordering-platform/pkg/utils/context"
	"ordering-platform/rpc/sys/client/menuservice"
	"ordering-platform/rpc/sys/sysclient"
	"sort"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserMenusLogic {
	return &UserMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserMenusLogic) UserMenus() (resp []*types.Menu, err error) {
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

	menuResult := convertListMenuDataToMenu(menus.Data)

	return menuResult, nil
}

func convertListMenuDataToMenu(list []*menuservice.ListMenuData) []*types.Menu {
	menuMap := make(map[int64]*types.Menu)
	var roots []*types.Menu

	// 转换每个ListMenuData到Menu，并存入map
	for _, ld := range list {
		if ld.MenuType == "F" {
			continue
		}
		menu := &types.Menu{
			MenuId:       ld.MenuId,
			ParentMenuId: ld.ParentMenuId,
			Name:         ld.Name,
			Path:         ld.Path,
			Component:    ld.Component,
			MenuMate: &types.MenuMate{
				Title:                    ld.Title,
				Order:                    int64(ld.Sort),
				Icon:                     ld.Icon,
				HideInMenu:               ld.HideInMenu,
				KeepAlive:                !ld.IgnoreKeepAlive,
				Link:                     getLink(ld.LinkFlag, ld.Link),
				MenuVisibleWithForbidden: !ld.Disabled,
			},
		}
		menuMap[menu.MenuId] = menu
	}

	// 构建父子关系
	for _, menu := range menuMap {
		parentID := menu.ParentMenuId
		if parentID == 0 {
			roots = append(roots, menu)
		} else if parent, exists := menuMap[parentID]; exists {
			parent.Children = append(parent.Children, menu)
		}
	}

	// 对所有节点排序
	sortMenus(roots)
	for _, menu := range menuMap {
		sortMenus(menu.Children)
	}

	return roots
}

// 辅助函数：根据LinkFlag决定是否返回链接
func getLink(linkFlag bool, link string) string {
	if linkFlag {
		return link
	}
	return ""
}

// 辅助函数：按Order排序菜单
func sortMenus(menus []*types.Menu) {
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].MenuMate.Order < menus[j].MenuMate.Order
	})
}

//func convertListMenuDataToMenu(list []*menuservice.ListMenuData) []types.Menu {
//	// 1. 创建映射表，存储 MenuID 到 *Menu 的映射（用于挂载子节点）
//	menuMap := make(map[int64]*types.Menu)
//	// 2. 存储所有菜单的原始值（值类型）
//	var menus []types.Menu
//
//	// 第一次遍历：转换 ListMenuData 到 Menu（值类型）
//	for _, item := range list {
//		if item.MenuType == "M" || item.MenuType == "C" {
//			meta := &types.MenuMate{
//				Title:                    item.Title,
//				Order:                    int64(item.Sort),
//				Icon:                     item.Icon,
//				HideInMenu:               item.HideInMenu,
//				KeepAlive:                !item.IgnoreKeepAlive,
//				MenuVisibleWithForbidden: !item.Disabled,
//				Link:                     item.Link,
//			}
//
//			menu := types.Menu{
//				MenuId:       item.MenuId,
//				ParentMenuId: item.ParentMenuId,
//				Name:         item.Name,
//				Path:         item.Path,
//				Component:    item.Component,
//				MenuMate:     meta,
//				Children:     []*types.Menu{}, // Children 仍为指针切片
//			}
//
//			// 将菜单存入切片，并通过索引获取其指针
//			menus = append(menus, menu)
//		}
//	}
//
//	// 为映射表赋值（指向 menus 中的元素）
//	var rootMenus []types.Menu
//	for i := range menus {
//		menu := &menus[i]
//		menuMap[menu.MenuId] = menu
//		if menu.ParentMenuId == 0 {
//			rootMenus = append(rootMenus, *menu) // 存储值类型
//		}
//	}
//
//	// 第二次遍历：构建树形结构（通过指针操作）
//	for i := range menus {
//		menu := &menus[i]
//		parentID := menu.ParentMenuId
//		if parentID != 0 {
//			if parent, ok := menuMap[parentID]; ok {
//				parent.Children = append(parent.Children, menu) // 挂载指针
//			}
//		}
//	}
//
//	return rootMenus
//}
