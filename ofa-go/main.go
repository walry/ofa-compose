package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"ofa/lib/db"
	"ofa/lib/log"
	"ofa/models"
	"ofa/routes"
	"os"
	"path/filepath"
)

var (
	ConfigPath string
	AppConfigFile string
)

//初始化项目配置
func loadConfig(){
	//命令行设置配置目录
	flag.StringVar(&ConfigPath,"c","conf","set your config path.")
	flag.Parse()

	AppConfigFile = filepath.Join(ConfigPath,"app.yml")
	fmt.Println("AppConfigFile:",AppConfigFile)
	if _,err := os.Stat(AppConfigFile); err != nil {
		fmt.Println(err.Error())
		fmt.Println(AppConfigFile," not found.")
		os.Exit(0)
	}
	//viper管理配置
	viper.SetConfigFile(AppConfigFile)
	var err error
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("parse config file: %s ---fail: %s",AppConfigFile,err))
	}
}
//配置日志
func logInit(){
	logConfFile := filepath.Join(ConfigPath,"log.yml")
	if logConfFile == "" {
		fmt.Println("Can`t find ",logConfFile)
		return
	}
	log.Init(logConfFile)
	viper.SetConfigFile(logConfFile)
}

//配置数据库连接
func dbInit(){
	dbConfigFile := filepath.Join(ConfigPath,"db.yml")
	if dbConfigFile == "" {
		fmt.Println("Can`t find ",dbConfigFile)
		return
	}
	db.Init(dbConfigFile)
	viper.SetConfigFile(dbConfigFile)
}
//启动服务
func ListenAndServe(g *gin.Engine)  {
	addr := fmt.Sprintf("0.0.0.0:%d",viper.GetInt("port"))
	log.Info("http server listen on:", addr)
	if err := http.ListenAndServe(addr,g); err != nil {
		log.Error("server start fail!!!  ERROR:", err)
		os.Exit(0)
	}
}

func main(){
	//加载配置
	loadConfig()

	//初始化日志模块
	logInit()
	defer func() { log.Sync() }()

	//初始化数据库
	dbInit()
	defer func() { db.Close() }()

	//模型初始化，一般新建表
	models.Init()

	//设置gin运行模式
	m := viper.GetString("runMode")
	gin.SetMode(m)
	//实例化gin
	g := gin.New()
	routes.Load(g)
	ListenAndServe(g)
}
