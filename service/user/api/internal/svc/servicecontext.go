package svc

import (
	"github.com/lixvyang/betxin-micro/service/topic/rpc/topicsrv"
	"github.com/lixvyang/betxin-micro/service/user/api/internal/config"
	"github.com/lixvyang/betxin-micro/service/user/api/internal/middleware"
	"github.com/lixvyang/betxin-micro/service/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Admin  rest.Middleware

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	UserModel model.UserModel

	TopicRPC topicsrv.Topicsrv

	AdminAuthenticate rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DNS)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
		TopicRPC:  topicsrv.NewTopicsrv(zrpc.MustNewClient(c.TopicRPC)),
		Admin:     middleware.NewAdminMiddleware().Handle,
	}
}
