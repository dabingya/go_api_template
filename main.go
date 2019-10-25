package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"learn_go/go_api_template/config"
	"learn_go/go_api_template/libs/db"
	"learn_go/go_api_template/router"
	"os"
	"path/filepath"
)

var (
	confFileName = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	// 初始化日志
	filePath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if err := config.InitLog(filePath); err != nil {
		panic(err)
		return
	}
	// 初始化配置，出错退出程序
	if err := config.InitConfig(*confFileName); err != nil {
		panic(err)
		return
	}
	g := gin.New()

	// 加载middleware
	router.Load(
		g,
	)
	//初始化数据库
	if err := db.InitDB(); err != nil {
		panic(err)
		return
	}

	_ = g.Run(viper.GetString("addr"))
}
