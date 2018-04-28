/**
 * Created by chaolinding on 2018/3/27.
 */

/*
mysql 提供mysql的增删改查等操作
 */
package mysql

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"config"
)

/*
mysql连接池
 */
var MysqlConn *sqlx.DB

func init() {
	InitMysqlPool()
}

/*
初始化mysql连接
 */
func InitMysqlPool() {
	var err error
	dataSourceName := config.GetMysqlConfig()

	MysqlConn, err = sqlx.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("mysql Connection err===", err)
	}
	MysqlConn.SetMaxOpenConns(35)
	MysqlConn.SetMaxIdleConns(30)
}
