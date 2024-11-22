package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ordering-platform/api/admin/internal/logic/sys/role"
	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"
)

// role详情
func RoleInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewRoleInfoLogic(r.Context(), svcCtx)
		resp, err := l.RoleInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
