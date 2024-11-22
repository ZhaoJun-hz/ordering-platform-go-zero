package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ordering-platform/api/admin/internal/logic/sys/role"
	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"
)

// role更新
func RoleUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewRoleUpdateLogic(r.Context(), svcCtx)
		resp, err := l.RoleUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
