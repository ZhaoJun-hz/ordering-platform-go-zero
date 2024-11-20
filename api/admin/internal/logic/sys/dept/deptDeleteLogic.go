package dept

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// dept删除
func NewDeptDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptDeleteLogic {
	return &DeptDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptDeleteLogic) DeptDelete(req *types.DeptDeleteReq) (resp *types.DeptDeleteResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.DeptService.DeptDelete(l.ctx, &sysclient.DeptDeleteReq{
		DeptId: req.DeptId,
	})
	if err != nil {
		return nil, err
	}
	return
}
