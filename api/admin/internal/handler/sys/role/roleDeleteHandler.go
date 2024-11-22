package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ordering-platform/api/admin/internal/logic/sys/role"
	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"
)

// role删除
func RoleDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewRoleDeleteLogic(r.Context(), svcCtx)
		resp, err := l.RoleDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
