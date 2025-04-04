package api

import (
	"bytes"
	"context"
	"github.com/bitly/go-simplejson"
	"ordering-platform/rpc/sys/sysclient"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"ordering-platform/pkg/global"
)

type InitApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 初始化api
func NewInitApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitApiLogic {
	return &InitApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitApiLogic) InitApi(req *types.InitApiReq) (resp *types.InitApiResp, err error) {
	// todo: add your logic here and delete this line

	routes := global.Routes
	list := make([]*sysclient.InitApiRouteData, 0)
	for _, route := range routes {
		dir, _ := os.Getwd()
		paths := filepath.Join(dir, "docs/admin.json")
		jsonFile, _ := os.ReadFile(paths)
		jsonData, _ := simplejson.NewFromReader(bytes.NewReader(jsonFile))
		if route.Method != "HEAD" {
			routeData := sysclient.InitApiRouteData{}
			routeType := "BUS"
			urlPath := route.Path
			idPatten := "(.*)/:(\\w+)" // 正则替换，把:id换成{id}
			reg, _ := regexp.Compile(idPatten)
			if reg.MatchString(urlPath) {
				urlPath = reg.ReplaceAllString(route.Path, "${1}/{${2}}") // 把:id换成{id}
			}
			if strings.HasPrefix(urlPath, "/api/sys/") {
				routeType = "SYS"
			}
			apiTitle := ""
			handler := ""
			apiTitle, _ = jsonData.Get("paths").Get(urlPath).Get(strings.ToLower(route.Method)).Get("summary").String()
			if apiTitle == "" {
				apiTitle, _ = jsonData.Get("paths").Get(urlPath + "/").Get(strings.ToLower(route.Method)).Get("summary").String()
			}
			handler, _ = jsonData.Get("paths").Get(urlPath).Get(strings.ToLower(route.Method)).Get("operationId").String()
			if handler == "" {
				handler, _ = jsonData.Get("paths").Get(urlPath + "/").Get(strings.ToLower(route.Method)).Get("operationId").String()
			}

			routeData.Title = apiTitle
			routeData.Path = urlPath
			routeData.Action = route.Method
			routeData.Handle = handler
			routeData.Type = routeType
			list = append(list, &routeData)
		}

	}

	_, err = l.svcCtx.ApiService.InitApi(l.ctx, &sysclient.InitApiReq{
		List: list,
	})
	if err != nil {
		return nil, err
	}
	return
}
