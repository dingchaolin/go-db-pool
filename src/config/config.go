/**
* Created by chaolinding on 2018/4/23
*/

package config

import (
	"utils"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var config *Config

type RedisConfig struct {
	Host     string
	Port     int
	Db       int
	Password string
}

type MongoConfig struct {
	Host     string
	Port     int
	UserName string
	Password string
	Database string
}
type MysqlConfig struct {
	Host   string
	Name   string
	Port   string
	User   string
	Passwd string
}
type Config struct {
	Mongo MongoConfig
	Mysql MysqlConfig
	Redis RedisConfig
}

var GetConfig = func() (*Config) {
	if config == nil {
		configPath := utils.RootDir + "/config/" + GetEnv() + "/config.yaml"

		var err error
		bytes, err := ioutil.ReadFile(configPath)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(bytes, &config)
		if err != nil {
			panic(err)
		}
	}
	return config
}



var GetMysqlConfig = func() (string) {
	return GetConfig().Mysql.User + ":" + GetConfig().Mysql.Passwd + "@tcp(" + GetConfig().Mysql.Host + ":" +
		GetConfig().Mysql.Port + ")/" + GetConfig().Mysql.Name + "?charset=utf8"
}

var GetMongoConfig = func() (MongoConfig) {
	return GetConfig().Mongo
}

var GetRedisConfig = func() RedisConfig {
	return GetConfig().Redis
}
