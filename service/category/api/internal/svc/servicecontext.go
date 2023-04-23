package svc

import (
	"github.com/lixvyang/betxin-micro/service/category/api/internal/config"
	"github.com/lixvyang/betxin-micro/service/category/api/internal/middleware"
	"github.com/lixvyang/betxin-micro/service/category/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	Admin  rest.Middleware
	CategoryModel model.CategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DNS)
	return &ServiceContext{
		Config: c,
		Admin:  middleware.NewAdminMiddleware().Handle,
		CategoryModel: model.NewCategoryModel(conn, c.CacheRedis),
	}
}
