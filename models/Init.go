package models

import (
	"fmt"
	"github.com/goBack/pkg/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "gorm_example"
)

var DB *gorm.DB

func Init() {
	util.Log().Info("initializing database...")
	var (
		db  *gorm.DB
		err error
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		util.Log().Panic("database connection failed...")

	}
	DB = db
	Migration()
}

func Migration() {
	//DB.DropTable(&Track{})
	DB.AutoMigrate(&Track{}, &LogButMon{})

}
