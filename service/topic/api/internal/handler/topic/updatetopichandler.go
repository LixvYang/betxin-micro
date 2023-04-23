package topic

import (
	"net/http"

	"github.com/lixvyang/betxin-micro/service/topic/api/internal/logic/topic"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateTopicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateTopicReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := topic.NewUpdateTopicLogic(r.Context(), svcCtx)
		resp, err := l.UpdateTopic(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
