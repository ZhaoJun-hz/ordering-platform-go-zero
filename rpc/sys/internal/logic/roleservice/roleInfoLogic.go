package roleservicelogic

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

	var selectMenuIds = make([]int64, 0)
	for _, roleMenu := range roleMenus {
		selectMenuIds = append(selectMenuIds, roleMenu.MenuID)
	}

	return &sysclient.RoleInfoResp{
		RoleId:        role.RoleID,
		RoleName:      role.RoleName,
		RoleKey:       role.RoleKey,
		Status:        role.Status,
		Sort:          role.RoleSort,
		SelectMenus:   selectMenuIds,
		DefaultRouter: role.DefaultRouter,
		Admin:         role.Admin,
	}, nil
}

type MenuTree struct {
	MenuId       int64       // 菜单ID
	ParentMenuId int64       // 父菜单ID
	Children     []*MenuTree // 子菜单
}
