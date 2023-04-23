package oauth

import (
	"net/http"

	"github.com/lixvyang/betxin-micro/service/user/api/internal/logic/oauth"
	"github.com/lixvyang/betxin-micro/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OauthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := oauth.NewOauthLogic(r.Context(), svcCtx, w, r)
		err := l.Oauth()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
