package main

import (
	"embed"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed dist/*
var FS embed.FS

func main() {
	// 从命令行接收端口参数
	var port string
	flag.StringVar(&port, "port", "9999", "HTTP服务器端口")
	flag.Usage = func() {
		fmt.Printf("使用方法 %s：\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	//创建fiber应用
	app := fiber.New()

	// 将嵌入的 dist 文件夹内容提供为静态文件服务
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(FS),
		PathPrefix: "dist",
		MaxAge:     3600, // 设置缓存控制头，缓存时间为1小时
	}))

	// 启动 HTTP 服务器
	go func() {
		log.Panic(app.Listen(":" + port))
	}()

	//保持运行
	select {}
}
