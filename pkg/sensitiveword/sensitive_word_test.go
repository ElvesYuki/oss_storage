package sensitiveword

import (
	"fmt"
	"go.uber.org/zap"
	"oss_storage/setting"
	"oss_storage/setting/logger"
	"oss_storage/setting/mysql"
	"testing"
)

func init() {
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
}

func TestSensitiveFilter(t *testing.T) {

}

func TestInsertSensitiveWord(t *testing.T) {

	// 打开文件
	/*file, err := os.OpenFile("../../static/sensi_words.bak.txt", os.O_RDONLY, 0600)
	if err != nil {
		zap.L().Error("打开敏感词文件出错", zap.Error(err))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	trieRoot = newTrieNode(' ')
	// 扫描每一行
	for scanner.Scan() {
		// 构建根节点树
		word := &sensitiveWord{Content: scanner.Text()}
		err := insertSensitiveWord(word)
		if err != nil {
			zap.L().Error("插入出错", zap.Error(err))
		}
	}*/
}
