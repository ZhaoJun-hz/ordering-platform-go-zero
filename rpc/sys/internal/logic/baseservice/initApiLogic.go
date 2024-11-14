package baseservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/gen/model"
	"ordering-platform/rpc/sys/gen/query"
	"time"

	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitApiLogic {
	return &InitApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 初始化Api
func (l *InitApiLogic) InitApi(in *sysclient.InitApiReq) (*sysclient.InitApiResp, error) {
	// todo: add your logic here and delete this line
	q := query.SysAPI
	for _, routeData := range in.List {
		_, err := q.WithContext(l.ctx).Where(q.Path.Eq(routeData.Path), q.Action.Eq(routeData.Action)).
			Attrs(field.Attrs(&model.SysAPI{
				Handle:    routeData.Handle,
				Title:     routeData.Title,
				Type:      routeData.Type,
				CreatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
				CreateBy:  1,
				UpdateBy:  1,
			})).FirstOrCreate()
		if err != nil {
			logc.Errorf(l.ctx, "初始化api异常,参数：%+v,异常:%s", in, err.Error())
			return nil, errors.Wrapf(xerr.NewDBErr(), "初始化api异常 %v, req %v", err, in)
		}
	}

	return &sysclient.InitApiResp{}, nil
}
