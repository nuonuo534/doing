package models

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"server/setting"
)

var sqlDB *sql.DB

type Model struct {
	ID         int `gorm: "primary_key" json:"id"`
	CreatedOn  int `json: "created_on"`
	ModifiedOn int `json: "modified_on"`
}

func init() {
	var (
		err error
		// dbType   string
		dbName   string
		user     string
		password string
		host     string
	)
	sec, ok := setting.Config["database"].(map[string]interface{})
	if !ok {
		log.Fatal(2, "Fail to get section database: %v", err)
	}

	// dbType = sec["type"].(string)
	dbName = sec["name"].(string)
	user = sec["user"].(string)
	password = sec["password"].(string)
	host = sec["host"].(string)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	sqlDB, err = db.DB()

	// db.SingularTable(true)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

}

func CloseDB() {
	defer sqlDB.Close()
}
