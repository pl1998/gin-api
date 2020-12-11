package redis

import (
	"github.com/go-redis/redis"
	"goproject/pkg/config"
	"goproject/pkg/log"
	"time"
)

var RedisDB *redis.Client

// redis 连接
func InitClient() (err error) {

	RedisDB = redis.NewClient(&redis.Options{
		Network:"tcp",
		Addr:config.GetString("cache.redis.addr")+":"+config.GetString("cache.redis.port"),
		Password:config.GetString("cache.redis.password"),
		DB:config.GetInt("cache.redis.db",0),
		PoolSize: 15,      //连接池 默认为4倍cpu数
		MinIdleConns: 10, //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量
		DialTimeout:5*time.Second,
		ReadTimeout:5*time.Second,
		WriteTimeout:5*time.Second,
		PoolTimeout:5*time.Second,
	})
	_,err = RedisDB.Ping().Result()

	if err != nil {
		log.LogError(err)
	}
	return nil
}