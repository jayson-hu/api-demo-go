package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jayson-hu/api-demo-go/apps/host"
)
var (
	API = &Handler{}
)

func NewHostHTTPHander(svc host.Service) *Handler  {
	return &Handler{
		svc: svc,
	}
}
//通过写一个实体类，来处理http协议，然后把内部的接口暴露出去
// 所以需要依赖内部接口的实现
// 该实体类会实现gin 的handler
type Handler struct {
	svc host.Service
}

//只是完成了http handlerd的注册
func (h *Handler) Register(r gin.IRouter)  {
	r.POST("/hosts", h.createHost)

}