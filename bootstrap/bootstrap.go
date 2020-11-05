package bootstrap

import (
	"context"
	"fmt"
	"gin-admin/app/utility/app"
	"gin-admin/app/utility/system"
	"gin-admin/database"
	"gin-admin/routers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
)

func init() {
	app.Panic(godotenv.Load())

	// 初始化所有工具类
	app.Panic(app.Redis().Connect())
	app.Panic(app.ConnectDB())

	// 自动创建数据表
	database.AutoGenTable()
	// 汉化参数验证器
	binding.Validator = new(system.Validator)
}

func Run() {
	gin.SetMode(os.Getenv("APP_ENV"))
	engine := gin.Default()
	// 设置路由
	routers.SetupRouter(engine)

	address := os.Getenv("HTTP_ADDRESS")
	port := os.Getenv("HTTP_PORT")
	endPoint := fmt.Sprintf("%s:%s", address, port)

	server := &http.Server{
		Addr:           endPoint,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Println(fmt.Sprintf("Server      Name:     %s", os.Getenv("APP_NAME")))
	fmt.Println(fmt.Sprintf("System      Name:     %s", runtime.GOOS))
	fmt.Println(fmt.Sprintf("Go          Version:  %s", runtime.Version()[2:]))
	fmt.Println(fmt.Sprintf("Gin         Version:  %s", gin.Version))
	fmt.Println(fmt.Sprintf("Listen      Address:  %s", "http://"+endPoint))

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("关闭服务...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("服务关闭错误:", err.Error())
	}
	log.Println("服务已关闭")
}
