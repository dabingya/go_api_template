package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	dbConn *sql.DB
	err    error
)

func InitDB() error {
	dbName := viper.GetString("db.name")
	dbUser := viper.GetString("db.user")
	dbPassword := viper.GetString("db.password")
	dbAddr := viper.GetString("db.addr")

	//dbConn, err = sql.Open("mysql", "oss_backup:0uOSsHoN6n3%GmDs@tcp(172.16.167.26:33306)/oss_backup?charset=utf8")
	dbConn, err = sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbAddr+")/"+dbName+"?charset=utf8")
	if err != nil {
		return err
	}
	if err := dbConn.Ping(); err != nil {
		logrus.Error("数据库连接失败" + err.Error())
		return err
	}
	dbConn.SetMaxOpenConns(60)
	dbConn.SetMaxIdleConns(100)
	logrus.Info("数据库连接成功")
	return nil
}
