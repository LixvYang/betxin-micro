package category

import (
	"net/http"

	"github.com/lixvyang/betxin-micro/service/category/api/internal/logic/category"
	"github.com/lixvyang/betxin-micro/service/category/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/category/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCategoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewGetCategoryLogic(r.Context(), svcCtx)
		resp, err := l.GetCategory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
