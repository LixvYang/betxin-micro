package svc

import (
	"github.com/lixvyang/betxin-micro/service/category/rpc/categorysrv"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/config"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	Admin rest.Middleware

	CategoryRPC categorysrv.Categorysrv

	// CategoryMap map[int64]*model.Category
}

func NewServiceContext(c config.Config) *ServiceContext {
	categoryRPC := categorysrv.NewCategorysrv(zrpc.MustNewClient(c.CategoryRPC))

	return &ServiceContext{
		Config:      c,
		Admin:       middleware.NewAdminMiddleware().Handle,
		CategoryRPC: categoryRPC,
	}
}
