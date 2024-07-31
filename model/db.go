package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"teacher2/utils"
)

var db *sql.DB
var err error

func InitDb() {
	db, err = sql.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPost,
		utils.DbName,
	))
	err = db.Ping()
	if err != nil {
		fmt.Printf("连接数据库失败，请检查参数:%s", err)
	}

	/*db.SingularTable(true)
	db.AutoMigrate(&User{}, &Admin{})*/
}
