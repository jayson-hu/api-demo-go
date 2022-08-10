package http

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
	"github.com/jayson-hu/api-demo-go/apps/host"
)

//用于暴露Host Service接口

func (h *Handler) createHost(c *gin.Context)  {
	ins := host.NewHost()
	//用户传递过来的函数进行解析
	if err := c.Bind(ins); err != nil {
		response.Failed(c.Writer, err)
		return
	}

	//进行接口调用
	ins, err := h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		return
	}
	response.Success(c.Writer, ins)

}
func (h *Handler) queryHost(c *gin.Context) {
	// 从http请求的query string 中获取参数
	req := host.NewQueryHostFromHTTP(c.Request)

	// 进行接口调用, 返回 肯定有成功或者失败
	set, err := h.svc.QueryHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, set)
}
func (h *Handler) DescribeHost(c *gin.Context) {
	// 从http请求的query string 中获取参数
	req := host.NewDescribeHostWithId(c.Param("id"))

	// 进行接口调用, 返回 肯定有成功或者失败
	set, err := h.svc.DescribeHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, set)
}