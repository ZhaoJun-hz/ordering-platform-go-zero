package userservicelogic

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

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// user列表
func (l *UserListLogic) UserList(in *sysclient.UserListReq) (*sysclient.UserListResp, error) {
	// todo: add your logic here and delete this line
	sysUser := query.SysUser
	q := sysUser.WithContext(l.ctx)
	if in.RoleId != 0 {
		q = q.Where(sysUser.RoleID.Eq(in.RoleId))
	}
	if in.DeptId != 0 {
		q = q.Where(sysUser.DeptID.Eq(in.DeptId))
	}
	if in.Username != "" {
		q = q.Where(sysUser.Username.Like("%" + in.Username + "%"))
	}
	if in.Nickname != "" {
		q = q.Where(sysUser.Nickname.Like("%" + in.Nickname + "%"))
	}
	if in.Email != "" {
		q = q.Where(sysUser.Email.Like("%" + in.Email + "%"))
	}
	if in.Mobile != "" {
		q = q.Where(sysUser.Mobile.Like("%" + in.Mobile + "%"))
	}

	offset := (in.PageNum - 1) * in.PageSize
	sysUsers, count, err := q.FindByPage(int(offset), int(in.PageSize))
	if err != nil {
		logc.Errorf(l.ctx, "查询 User 列表失败, 异常:%s", err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 User 列表失败 %v", err)
	}

	var result []*sysclient.UserListData
	for _, item := range sysUsers {
		result = append(result, &sysclient.UserListData{
			UserId:      item.UserID,
			Status:      item.Status,
			Username:    item.Username,
			Nickname:    item.Nickname,
			Description: item.Description,
			Mobile:      item.Mobile,
			Email:       item.Email,
			Avatar:      item.Avatar,
			DeptId:      item.DeptID,
			RoleId:      item.RoleID,
			CreateTime:  utils.TimeToString(&item.CreatedAt),
		})
	}

	return &sysclient.UserListResp{
		Total: count,
		Data:  result,
	}, nil
}
