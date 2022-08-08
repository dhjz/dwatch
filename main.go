package main

import (
	"dwatch/common"
	"dwatch/routers"
	"dwatch/schedule"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	port := flag.Int("p", 3457, "server port")
	isStop := flag.Int("s", 0, "stop watch")
	flag.Parse()

	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()
	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//创建一个默认的路由引擎
	c := gin.Default()
	c.Use(common.Cors())
	//添加路由
	routers.RegisterRouter(c)
	addr := fmt.Sprintf(":%d", *port)
	// addr := ":" + strconv.Itoa(*port)
	schedule.WatchAll(*isStop)
	fmt.Printf("http server start at port:"+addr, "stop watch:", isStop)
	c.Run(addr) // 监听并在 0.0.0.0:8080 上启动服务

}
