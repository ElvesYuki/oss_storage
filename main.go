package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"oss_storage/pkg/idgenerator"
	"oss_storage/route"
	"oss_storage/setting"
	"oss_storage/setting/logger"
	"oss_storage/setting/mysql"
	"oss_storage/setting/oss"
	"oss_storage/setting/redis"
	"syscall"
	"time"
)

// @title Oss对象存储
// @version 1.0
// @description Oss对象存储
// @termsOfService https://elvesyuki.com

// @contact.name luohuan
// @contact.url https://elvesyuki.com
// @contact.email 1026770043@qq.com

// @host localhost:9091
// @BasePath
func main() {
	// 1、加载 配置
	fmt.Println("1、加载 配置")
	err := setting.Init()
	if err != nil {
		fmt.Printf("init settting failed, err:%v\n", err)
		return
	}

	// 2、初始化 日志
	fmt.Println("2、初始化 日志")
	err = logger.Init(setting.Conf.LogConfig, setting.Conf.Mode)
	if err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()

	// 3、初始化 MySQL 的连接
	fmt.Println("3、初始化 MySQL 的连接")
	err = mysql.Init(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	// 4、初始化 Redis 的连接
	fmt.Println("4、初始化 Redis 的连接")
	err = redis.Init(setting.Conf.RedisConfig)
	if err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	// 5、注册路由
	fmt.Println("5、注册路由")
	r := route.SetUp()

	// 6、初始化id生成器
	fmt.Println("6、初始化id生成器")
	err = idgenerator.Init()
	if err != nil {
		fmt.Printf("init idgenerator failed, err:%v\n", err)
		return
	}

	// 7、初始化MinioClient
	fmt.Println("7、初始化MinioClient")
	err = oss.Init(setting.Conf.OssConfig)
	if err != nil {
		fmt.Printf("init minio client failed, err:%v\n", err)
		return
	}

	// 启动服务(优雅关机)
	fmt.Println("6、启动服务(优雅关机)")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Conf.AppDetailConfig.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen: %s\n", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
