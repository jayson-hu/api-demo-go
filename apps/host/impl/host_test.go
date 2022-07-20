package impl_test

import (
	"context"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/jayson-hu/api-demo-go/apps/host"
	"github.com/jayson-hu/api-demo-go/apps/host/impl"
	"testing"
)

var  (
	//定义的对象是满足该接口的实例，
	service host.Service
)

func TestCreate(*testing.T)  {
	ins := host.NewHost()
	ins.Name="test"
	service.CreateHost(context.Background(), ins)


}
func init()  {
	//初始化logger
	//为什么不设计默认打印.因为性能
	zap.DevelopmentSetup()
	//host service 的具体实现
	service = impl.NewHosServiceImpl()



}