package common

import "go.uber.org/zap"

func ManualRecover() {
	if err := recover(); err != nil {
		zap.L().Error("手动捕获到panic")
	}
}
