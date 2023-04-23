package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/lixvyang/betxin-micro/service/topic/rpc/topicsrv"
	"github.com/lixvyang/betxin-micro/service/user/api/internal/config"
	"github.com/lixvyang/betxin-micro/service/user/model"
)

type ServiceContext struct {
	Config config.Config

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	UserModel model.UserModel

	TopicRPC topicsrv.Topicsrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DNS)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
		TopicRPC:  topicsrv.NewTopicsrv(zrpc.MustNewClient(c.TopicRPC)),
	}
}
