package svc

import (
	"github.com/lixvyang/betxin-micro/service/category/model"
	"github.com/lixvyang/betxin-micro/service/category/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	CategoryModel model.CategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DNS)
	return &ServiceContext{
		Config:        c,
		CategoryModel: model.NewCategoryModel(conn, c.CacheRedis),
	}
}
