package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jayson-hu/api-demo-go/apps/host"
)

//IOC容器层： 管理所有服务的实例

// 1.HostService的实例，必须注册过来
// 2. Http暴露模块，依赖Ioc 里面的HostService

var (
	HostService host.Service //
	implApps    = map[string]ImplService{}
	ginApps     = map[string]GinService{}
)

func RegistryImpl(svc ImplService) {

	//服务注册到svc的map中
	if _, ok := implApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registry", svc.Name()))
	}
	implApps[svc.Name()] = svc
	//根据满足的接口来注册具体的服务
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
	//switch obj.(type) {
	//case host.Service:
	//
	//
	//}
}

//如果指定了具体类型，就导致每增加一种类型，就多个GET方法
//func GetHostImpl(name string) host.Service

//get 一个impl服务的实例 implApps
//返回任何对象，任何对象都可以了，使用时，再使用方法进行断言
func GetImpl(name string) interface{} {
	for k, v := range implApps {
		if k == name {
			return v
		}
	}
	return nil
}

func RegistryGin(svc GinService) {
	if _, ok := ginApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registry", svc.Name()))
	}
	ginApps[svc.Name()] = svc
}

//用户初始化 注册到IOC容器里面的所有服务

func InitImpl() {
	for _, v := range implApps {
		v.Config()
	}
}

// LoadedGinApps 已经加载完成的gin app有哪些
func LoadedGinApps() (names []string) {
	for k := range ginApps {
		names = append(names, k)
	}
	return names
}

func InitGin(r gin.IRouter) {
	//把所有对象进行初始化
	for _, v := range ginApps {
		v.Config()
	}
	//完成 http handler 的注册
	for _, v := range ginApps {
		v.Registry(r)
	}
}

type ImplService interface {
	Config()
	Name() string
}

//注册由gin来注册
// 比如编写了http服务A，只需要实现Registry方法,就能把handler注册给 root Router

type GinService interface {
	Registry(r gin.IRouter)
	Config()
	Name() string
}
