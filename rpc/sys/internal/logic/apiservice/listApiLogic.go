package apiservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"ordering-platform/pkg/common"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/gen/query"

	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListApiLogic {
	return &ListApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListApiLogic) ListApi(in *sysclient.ApiListReq) (*sysclient.ApiListResp, error) {
	// todo: add your logic here and delete this line
	sysAPI := query.SysAPI
	q := sysAPI.WithContext(l.ctx)
	if in.Action != "" {
		q = q.Where(sysAPI.Action.Eq(in.Action))
	}
	if in.Type != "" {
		q = q.Where(sysAPI.Type.Eq(in.Type))
	}
	offset := (in.PageNum - 1) * in.PageSize
	result, count, err := q.FindByPage(int(offset), int(in.PageSize))
	if err != nil {
		logc.Errorf(l.ctx, "查询Api List 异常,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询Api List 异常 %v, req %v", err, in)
	}

	var list []*sysclient.ApiInfo
	for _, item := range result {
		list = append(list, &sysclient.ApiInfo{
			Id:         item.ID,
			Handle:     item.Handle,
			Title:      item.Title,
			Path:       item.Path,
			Type:       item.Type,
			Action:     item.Action,
			CreateTime: common.TimeToString(&item.CreatedAt),
		})
	}

	return &sysclient.ApiListResp{
		Total: count,
		Data:  list,
	}, nil
}
