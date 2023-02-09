package mysql

import (
	"database/sql"
	"fmt"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	log "github.com/sirupsen/logrus"
	"go-chatgpt-api/config"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Init 注册mysql
func Init() {
	log.Println("connect mysql start~")
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		beego.Error("mysql register driver error:", err)
	}

	username := config.GetMysqlConfig().Username
	password := config.GetMysqlConfig().Password
	host := config.GetMysqlConfig().Host
	port := config.GetMysqlConfig().Port
	database := config.GetMysqlConfig().Database

	createDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci", database)
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(createDB)
	if err != nil {
		panic(err)
	}

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", username, password, host, port, database)

	err = orm.RegisterDataBase("default", "mysql", dataSource)
	if err != nil {
		beego.Error("mysql register database error:", err)
	}
	//每次自动创建表，会清除表数据
	//orm.RunSyncdb("default", true, true)
	log.Println("connect mysql success~")
}
