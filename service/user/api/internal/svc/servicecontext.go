package svc

import (
	"github.com/lixvyang/betxin-micro/service/user/api/internal/config"
	"github.com/lixvyang/betxin-micro/service/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DNS)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
