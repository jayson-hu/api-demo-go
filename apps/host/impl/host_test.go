package impl_test

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/jayson-hu/api-demo-go/apps/host"
	"github.com/jayson-hu/api-demo-go/apps/host/impl"
	"github.com/jayson-hu/api-demo-go/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	//定义的对象是满足该接口的实例，
	service host.Service
)

func TestCreate(t *testing.T) {
	should := assert.New(t)
	ins := host.NewHost()
	ins.Id = "test-01"
	ins.Name = "test"
	ins.Region = "hangzhou"
	ins.Type = "1"
	ins.CPU = 2
	ins.Memory = 2
	ins, err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}

}

func TestQuery(t *testing.T) {
	should := assert.New(t)

	req := host.NewQueryHostRequest()
	req.Keywords = "test"
	set, err := service.QueryHost(context.Background(), req)
	if should.NoError(err) {
		for i := range set.Items {
			fmt.Println(set.Items[i].Id)
		}
	}
}

func TestDescribe(t *testing.T) {
	should := assert.New(t)

	req := host.NewDescribeHostWithId("test-01")
	ins, err := service.DescribeHost(context.Background(), req)
	if should.NoError(err) {
		fmt.Println(ins.Name)
	}
}


func init() {
	err := conf.LoadConfigFromToml("D:\\GoProject\\go-course-demo\\api-demo-go\\etc\\demo.toml")
	//err := conf.LoadConfigFromTEnv()
	if err != nil {
		panic(err)
	}
	//初始化logger
	//为什么不设计默认打印.因为性能
	zap.DevelopmentSetup()
	//host service 的具体实现
	service = impl.NewHostServiceImpl()

}
