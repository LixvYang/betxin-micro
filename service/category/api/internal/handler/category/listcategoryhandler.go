package category

import (
	"net/http"

	"github.com/lixvyang/betxin-micro/service/category/api/internal/logic/category"
	"github.com/lixvyang/betxin-micro/service/category/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := category.NewListCategoryLogic(r.Context(), svcCtx)
		resp, err := l.ListCategory()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
