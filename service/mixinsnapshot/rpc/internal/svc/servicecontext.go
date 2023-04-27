package svc

import "github.com/lixvyang/betxin-micro/service/mixinsnapshot/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
