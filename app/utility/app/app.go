package app

import (
	"gin-admin/app/utility/mysql"
	"gin-admin/app/utility/redis"
	"gin-admin/config"
	"go.uber.org/dig"
	"net/http"
	"xorm.io/xorm"
)

var container *dig.Container

func init() {
	container = dig.New()
	// 注入配置文件
	_ = container.Provide(config.New)
	// 注入redis组件
	_ = container.Provide(newRedis)
	// 注入mysql组件
	_ = container.Provide(newDBManager)
	// 注入http客户端
	_ = container.Provide(newHttpClient)
}

func Redis() (client *redis.Client) {
	_ = container.Invoke(func(cli *redis.Client) {
		client = cli
	})
	return
}

func dbManager() (m *mysql.Manager) {
	_ = container.Invoke(func(instance *mysql.Manager) {
		m = instance
	})
	return
}

func ConnectDB() error {
	return dbManager().Connect()
}

// DB 返回数据库连接
func DB() (conn *xorm.Engine) {
	return dbManager().Engine
}

// Http client
func Http() (client *http.Client) {
	_ = container.Invoke(func(instance *http.Client) {
		client = instance
	})
	return
}