package driverx

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Conn struct {
	Mysql sqlx.SqlConn
	Redis *redis.Redis
}

type Config struct {
	DataSource string
	CacheRedis cache.ClusterConf
}

func InitDriver(c *Config) (*Conn, error) {
	return &Conn{
		Mysql: sqlx.NewMysql(c.DataSource),
		Redis: redis.MustNewRedis(c.CacheRedis[0].RedisConf),
	}, nil
}
