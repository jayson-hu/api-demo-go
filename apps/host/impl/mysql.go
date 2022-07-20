package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/jayson-hu/api-demo-go/apps/host"
)

//接口实现的静态检查
var _ host.Service = (*HostServiceImpl)(nil)

func NewHosServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		// host service 服务的 Newlogger
		// 封装的zap 让其满足 Logger接口
		//为什么要封装
			// 1. Logger 全局实例 2.logger level的动态配置调整 logrus不支持level共同调整 3. 加入日志轮转的功能的集合
		l: zap.L().Named("Host"),
	}
}

type HostServiceImpl struct {
	l logger.Logger
}
