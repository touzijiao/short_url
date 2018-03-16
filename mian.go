package main

import (
	"short_url/controllers"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	port string
)

func convertLogLevel(level string) int {

	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}

	return logs.LevelDebug
}

func init() {
	config := make(map[string]interface{}) //日志属性写到map内
	config["filename"] = "./logs/Short_Url.log"
	config["level"] = convertLogLevel("debug")

	configJSON, err := json.Marshal(config) //序列化
	if err != nil {
		fmt.Println("initLogger failed, err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configJSON)) //把map配置进去
	logs.EnableFuncCallDepth(true)                       //true日志文件输出文件名与行号，false关闭日志文件输出文件名与行号
	return
}

func main() {
	flag.StringVar(&port, "port", ":8081", "port to listen")
	flag.Parse()

	fmt.Println("MongoHost=>", beego.AppConfig.String("mongo_host"))
	fmt.Println("MongoPort=>", beego.AppConfig.String("mongo_port"))

	/*
		RESTRouter与Router功能一样
	*/
	//请求跳转功能
	beego.RESTRouter("/api/url", &controllers.UrlController{}) //这个一般放在routers文件夹下，在init时就调用
	beego.Router("/*", &controllers.RedirectController{})      //端口上的任意请求

	//beego处理静态文件
	beego.SetStaticPath("/public", "views") //(路由路径，静态文件路径)

	beego.Run(port)
	//beego.Run()
}
