/**
 * Created by chaolinding on 2018/4/28.
 */

package main

import (
	"db/redis"
	"log"
	"config"
	"db/mongo"
)

func main(){
	c, err := redis_client.NewClient("127.0.0.1", 6379, "", 0)
	log.Println( "redis===", c, err )

	cfg := config.MongoConfig{
		Host:     "127.0.0.1",
		Port:     27017,
		UserName: "",
		Password: "",
		Database: "mytest",
	}
	client, err := mongo.NewClient(cfg)

	if err != nil {
		log.Println( err )
	}

	type Person struct {
		Names  string
	}

	err = client.C("mes").Insert(Person{Names:"ggg"})
	if err != nil {
		log.Println(err)
	}else{
		log.Println("插入成功")
	}
}
