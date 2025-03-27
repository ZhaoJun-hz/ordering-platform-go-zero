package roleservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"ordering-platform/pkg/utils"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/gen/query"

	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleListLogic) RoleList(in *sysclient.RoleListReq) (*sysclient.RoleListResp, error) {
	// todo: add your logic here and delete this line
	sysRole := query.SysRole
	q := sysRole.WithContext(l.ctx)
	if in.RoleName != "" {
		q = q.Where(sysRole.RoleName.Like("%" + in.RoleName + "%"))
	}
	if in.RoleKey != "" {
		q = q.Where(sysRole.RoleKey.Like("%" + in.RoleKey + "%"))
	}
	if in.Status != 0 {
		q = q.Where(sysRole.Status.Eq(in.Status))
	}
	offset := (in.PageNum - 1) * in.PageSize
	sysRoles, count, err := q.FindByPage(int(offset), int(in.PageSize))
	if err != nil {
		logc.Errorf(l.ctx, "查询 Role 列表失败, 异常:%s", err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Role 列表失败 %v", err)
	}

	var result []*sysclient.RoleListData
	for _, sysRole := range sysRoles {
		result = append(result, &sysclient.RoleListData{
			RoleId:        sysRole.RoleID,
			RoleName:      sysRole.RoleName,
			RoleKey:       sysRole.RoleKey,
			Status:        sysRole.Status,
			Sort:          sysRole.RoleSort,
			DefaultRouter: sysRole.DefaultRouter,
			CreateTime:    utils.TimeToString(&sysRole.CreatedAt),
			Admin:         sysRole.Admin,
		})
	}

	return &sysclient.RoleListResp{
		Total: count,
		Data:  result,
	}, nil
}
