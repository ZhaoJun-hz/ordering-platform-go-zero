package dept

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// dept更新
func NewDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptUpdateLogic {
	return &DeptUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptUpdateLogic) DeptUpdate(req *types.DeptUpdateReq) (resp *types.DeptUpdateResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.DeptService.DeptUpdate(l.ctx, &sysclient.DeptUpdateReq{
		DeptId:       req.DeptId,
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
