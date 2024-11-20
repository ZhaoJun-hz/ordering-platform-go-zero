package dept

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// dept详情
func NewDeptInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptInfoLogic {
	return &DeptInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptInfoLogic) DeptInfo(req *types.DeptInfoReq) (resp *types.DeptInfoData, err error) {
	// todo: add your logic here and delete this line
	info, err := l.svcCtx.DeptService.DeptInfo(l.ctx, &sysclient.DeptInfoReq{
		DeptId: req.DeptId,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeptInfoData{
		DeptId:       info.DeptId,
		ParentDeptId: info.ParentDeptId,
		DeptName:     info.DeptName,
		Sort:         info.Sort,
		Leader:       info.Leader,
		Phone:        info.Phone,
		Email:        info.Email,
		Status:       info.Status,
		CreateTime:   info.CreateTime,
	}, nil
}
