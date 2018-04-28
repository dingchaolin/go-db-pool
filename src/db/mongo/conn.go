/**
 * Created by chaolinding on 2018/4/28.
 */

package mongo

import (
	"gopkg.in/mgo.v2"
	"config"
	"fmt"
	"log"
	"time"
)

/*
mongo pool
记录feed信息
 */
type FeedSchema struct {
	Uid      int       `bson:"uid"`
	FId      int       `bson:"fid"`
	Text     string    `bson:"text"`
	CreateAt time.Time `bson:"createAt"`
}

var MongoPool *mgo.Database

func init() {
	InitMongoPool()
}

/*
mongo 连接初始化
 */
func InitMongoPool() {

	config := config.GetMongoConfig()
	var err error

	MongoPool, err = NewClient(config)
	if err != nil {
		log.Fatal("初始化mongo 失败！")
	}

}

func NewClient(config config.MongoConfig) (*mgo.Database, error) {

	url := ""

	if config.UserName != "" && config.Password != "" {

		url = fmt.Sprintf(
			"mongodb://%s:%s@%s:%d",
			config.UserName,
			config.Password,
			config.Host,
			config.Port)

	} else {

		url = fmt.Sprintf(
			"mongodb://%s:%d",
			config.Host,
			config.Port)
	}

	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	db := session.DB(config.Database)

	return db, err
}
