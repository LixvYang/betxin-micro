package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	Mysql struct {
		DNS string
	}

	CacheRedis cache.CacheConf
	Mixin      struct {
		Pin        string
		ClientId   string
		SessionId  string
		PinToken   string
		PrivateKey string
		AppSecret  string
	}

	AppMode string
}
