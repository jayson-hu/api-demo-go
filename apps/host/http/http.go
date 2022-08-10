package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jayson-hu/api-demo-go/apps"
	"github.com/jayson-hu/api-demo-go/apps/host"
)

var (
	API = &Handler{}
)

var handler = &Handler{}

func NewHostHTTPHander() *Handler {
	return &Handler{

	}
}

// Handler 通过写一个实体类，来处理http协议，然后把内部的接口暴露出去
// 所以需要依赖内部接口的实现
// 该实体类会实现gin 的handler
type Handler struct {
	svc host.Service
}

func (h *Handler) Config() {

	//if apps.HostService == nil {
	//	panic("dependency hsot service required")
	//}
	//从IOC 里面获取HostService 获取实例对象
	h.svc = apps.GetImpl(host.AppName).(host.Service)
}

//只是完成了http handlerd的注册
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
	r.GET("/hosts", h.queryHost)

}
func (h *Handler) Name() string {
	return host.AppName
}

//完成 http handler的注册
func init() {
	apps.RegistryGin(handler)
}
