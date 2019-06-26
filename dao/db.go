package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hardbornman/garglewool-service/initials/config"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var (
	garglewool *sqlx.DB
	err        error
)

func Init() {
	initGarglewool()
	initFields(config.Config.DB.DbName)
}

func initGarglewool() {
	garglewool, err = sqlx.Open("mysql", fmt.Sprintf(config.Config.DB.Dsn, config.Config.DB.Pwd))
	if err != nil {
		log.Fatalf("【initGarglewool.NewEngine】ex:%s", err.Error())
		os.Exit(0)
		return
	}
	err = garglewool.Ping()
	if err != nil {
		log.Fatalf("【initGarglewool.Ping】ex:%s", err.Error())
		os.Exit(0)
		return
	}
	garglewool.SetMaxIdleConns(config.Config.DB.MaxIdleConn)
	garglewool.SetMaxOpenConns(config.Config.DB.MaxOpenConn)
}
