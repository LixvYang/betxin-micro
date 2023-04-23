package topic

import (
	"net/http"

	"github.com/lixvyang/betxin-micro/service/topic/api/internal/logic/topic"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListTopicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := topic.NewListTopicLogic(r.Context(), svcCtx)
		resp, err := l.ListTopic()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
