package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	Mysql struct {
		Datasource string
	}

	Redis redis.RedisConf

	//系统
	SysRpc zrpc.RpcClientConf
}
