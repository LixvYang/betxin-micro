package svc

import (
	"github.com/lixvyang/betxin-micro/service/category/rpc/categorysrv"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	CategoryRPC categorysrv.Categorysrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		CategoryRPC: categorysrv.NewCategorysrv(zrpc.MustNewClient(c.CategoryRPC)),
	}
}
