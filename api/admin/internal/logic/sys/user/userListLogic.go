package user

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// user列表
func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.UserService.UserList(l.ctx, &sysclient.UserListReq{
		Username: req.Username,
		Nickname: req.Nickname,
		Mobile:   req.Mobile,
		Email:    req.Email,
		DeptId:   req.DeptId,
		RoleId:   req.RoleId,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var list []*types.UserListData
	for _, item := range result.Data {
		list = append(list, &types.UserListData{
			UserId:     item.UserId,
			Username:   item.Username,
			Nickname:   item.Nickname,
			Mobile:     item.Mobile,
			Email:      item.Email,
			Status:     item.Status,
			CreateTime: item.CreateTime,
		})
	}
	return &types.UserListResp{
		Total: result.Total,
		Data:  list,
	}, nil
	return
}
