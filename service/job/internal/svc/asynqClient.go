package svc

import (
	"github.com/hibiken/asynq"
	"github.com/lixvyang/betxin-micro/service/job/internal/config"
)

func newAsynqClient(c config.Config) *asynq.Client {
	return asynq.NewClient(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
	)
}
