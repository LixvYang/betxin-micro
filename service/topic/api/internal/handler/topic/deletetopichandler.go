package topic

import (
	"net/http"

	"github.com/lixvyang/betxin-micro/service/topic/api/internal/logic/topic"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteTopicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteTopicReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := topic.NewDeleteTopicLogic(r.Context(), svcCtx)
		resp, err := l.DeleteTopic(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
