/**
 * Created by chaolinding on 2018/4/23.
 */

package redis_client

import (
	"github.com/go-redis/redis"
	"strconv"
	"log"
	"config"
)

var redisMMJ *redis.Client
var redisDataSourceList []*redis.Client

func init() {
	InitRedisPool()
}

/*
redis 连接初始化
 */
func InitRedisPool() {

	mmjCfg := config.GetRedisConfig()
	var err error
	redisMMJ, err = NewClient(mmjCfg.Host, mmjCfg.Port, mmjCfg.Password, mmjCfg.Db)
	if err != nil {
		log.Fatal("初始化mmj redis pool 失败！")
	}

}

/*
创建一个client
 */
func NewClient(ip string, port int, password string, db int) (*redis.Client, error) {

	options := &redis.Options{
		Addr:      ip + ":" + strconv.Itoa(port),
		Password:  password,
		DB:        db,
	}
	client := redis.NewClient(options)

	_, err := client.Ping().Result()

	return client, err

}

func GetUserDataSource(uid int)*redis.Client{
	return redisDataSourceList[uid % len(redisDataSourceList)]
}

func GetRedis()*redis.Client{
	return redisMMJ
}

