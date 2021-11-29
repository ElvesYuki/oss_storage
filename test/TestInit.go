package test

import (
	"fmt"
	"go.uber.org/zap"
	"oss_storage/pkg/idgenerator"
	"oss_storage/setting"
	"oss_storage/setting/logger"
	"oss_storage/setting/mysql"
	"oss_storage/setting/oss"
	"oss_storage/setting/redis"
)

func InitTest() {
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
	//defer mysql.Close()

	// 4、初始化 Redis 的连接
	fmt.Println("4、初始化 Redis 的连接")
	err = redis.Init(setting.Conf.RedisConfig)
	if err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	//defer redis.Close()

	// 5、注册路由
	//fmt.Println("5、注册路由")
	//r := route.SetUp()

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

	// 8、初始化Sensitiveword
	//fmt.Println("8、初始化Sensitiveword")
	//sensitiveword.Init()
	//if err != nil {
	//	fmt.Printf("init Sensitiveword client failed, err:%v\n", err)
	//	return
	//}
}
