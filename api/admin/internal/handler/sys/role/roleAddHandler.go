package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ordering-platform/api/admin/internal/logic/sys/role"
	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"
)

// role添加
func RoleAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewRoleAddLogic(r.Context(), svcCtx)
		resp, err := l.RoleAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
