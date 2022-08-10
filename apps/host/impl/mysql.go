package impl

import (
	"database/sql"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/jayson-hu/api-demo-go/apps"
	"github.com/jayson-hu/api-demo-go/apps/host"
	"github.com/jayson-hu/api-demo-go/conf"
)

//接口实现的静态检查
//var _ host.Service = (*HostServiceImpl)(nil)

//这样写，会造成conf.c()并没有准备好, 造成方法panic
//var impl = NewHostServiceImpl()
//只有准备好对象，把对象的注册和初始化，独立出来
var impl = &HostServiceImpl{}

//NewHostServiceImpl 保证调用该函数之前已经完成全局config对象完成初始化
func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		// host service 服务的 Newlogger
		// 封装的zap 让其满足 Logger接口
		//为什么要封装
			// 1. Logger 全局实例 2.logger level的动态配置调整 logrus不支持level共同调整 3. 加入日志轮转的功能的集合
		l: zap.L().Named("Host"),
		db: conf.C().MySQL.GetDB(),
	}
}

type HostServiceImpl struct {
	l logger.Logger
	db *sql.DB
}

// 需要保证 全局对象config 和全局的logger已经加载完成
func (i *HostServiceImpl) Config()  {
	i.l = zap.L().Named("Host")
	i.db = conf.C().MySQL.GetDB()
}

// 返回服务的名称
func (i *HostServiceImpl) Name() string {
	return host.AppName
}
//  _ import app 自行注册对象
func init() {
	apps.RegistryImpl(impl)
	//apps.HostService = impl
}


//之前都是start的时候，手动把服务注册到IOC
//注册到HostService的实例到IOC
// apps.HostService= impl.NewHostServiceImpl

//mysql的驱动加载的实现方式
//sql 这个库，是一个框架，驱动是引用依赖的时候加载的
//我们吧app的模块，比作一个驱动，IOC比作框架



















