package menuservicelogic

import (
	"context"
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

type DeleteMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMenuLogic) DeleteMenu(in *sysclient.DeleteMenuReq) (*sysclient.DeleteMenuResp, error) {
	// todo: add your logic here and delete this line

	var err error
	_, err = query.SysMenu.WithContext(l.ctx).Where(query.SysMenu.MenuID.Eq(in.MenuId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "Menu不存在,参数：%d,异常:%s", in.MenuId, err.Error())
		return nil, errors.WithStack(errcode.MenuNotExistError)
	}
	if err != nil {
		logc.Errorf(l.ctx, "查询 Menu失败,参数：%d,异常:%s", in.MenuId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Menu失败 %v, param %v", err, in.MenuId)
	}
	// 校验一下有没有子菜单
	subMenuNumber, err := query.SysMenu.WithContext(l.ctx).Where(query.SysMenu.ParentMenuID.Eq(in.MenuId)).Count()
	if err != nil {
		logc.Errorf(l.ctx, "查询 Menu失败,参数：%d,异常:%s", in.MenuId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Menu失败 %v, param %v", err, in.MenuId)
	}
	if subMenuNumber > 0 {
		logc.Errorf(l.ctx, "Menu 下有子菜单,禁止删除")
		return nil, errors.WithStack(errcode.MenuHaveSubMenuError)
	}
	// 校验一下有没有角色关联该菜单，有的话不允许修改
	count, err := query.SysRoleMenu.WithContext(l.ctx).Where(query.SysRoleMenu.MenuID.Eq(in.MenuId)).Count()
	if err != nil {
		logc.Errorf(l.ctx, "查询 Menu 是否关联角色 失败,参数：%d,异常:%s", in.MenuId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Menu 是否关联角色 失败 %v, param %v", err, in.MenuId)
	}
	if count > 0 {
		logc.Errorf(l.ctx, "Menu 关联的有角色,不允许操作")
		return nil, errors.WithStack(errcode.MenuHaveAllocationRoleError)
	}

	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 删除sys_menu
	_, err = tx.SysMenu.WithContext(l.ctx).Where(query.SysMenu.MenuID.Eq(in.MenuId)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除 sys_menu失败,参数：%d,异常:%s", in.MenuId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "删除 sys_menu 失败 %v, 参数 %d", err, in.MenuId)
	}

	// 删除 sys_menu_api_rule
	_, err = tx.SysMenuAPI.WithContext(l.ctx).Where(query.SysMenuAPI.SysMenuID.Eq(in.MenuId)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除 sys_menu_api_rule 失败,参数：%d,异常:%s", in.MenuId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "删除 sys_menu_api_rule 失败 %v, 参数 %d", err, in.MenuId)
	}

	// TODO casbin 加一个字段 SysMenuAPIRule.id 方便删除
	// TODO 重新生成或者加载 casbin

	return &sysclient.DeleteMenuResp{}, nil
}
