package svc

import (
	"github.com/lixvyang/betxin-micro/service/topic/model"
	"github.com/lixvyang/betxin-micro/service/topic/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	TopicModel model.TopicModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:    c,
		TopicModel: model.NewTopicModel(conn, c.CacheRedis),
	}
}
