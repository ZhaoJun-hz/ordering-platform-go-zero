package dept

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// dept添加
func NewDeptAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptAddLogic {
	return &DeptAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptAddLogic) DeptAdd(req *types.DeptAddReq) (resp *types.DeptAddResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.DeptService.DeptAdd(l.ctx, &sysclient.DeptAddReq{
		ParentDeptId: req.ParentDeptId,
		DeptName:     req.DeptName,
		Sort:         req.Sort,
		Leader:       req.Leader,
		Phone:        req.Phone,
		Email:        req.Email,
		Status:       req.Status,
	})
	if err != nil {
		return nil, err
	}
	return
}
