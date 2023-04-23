package user

import (
	"net/http"

	"github.com/lixvyang/betxin-micro/service/user/api/internal/logic/user"
	"github.com/lixvyang/betxin-micro/service/user/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewCreateUserLogic(r.Context(), svcCtx)
		resp, err := l.CreateUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
