package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/jayson-hu/api-demo-go/apps"
	"github.com/jayson-hu/api-demo-go/conf"

)

func NewHttpService() *HttpService {
	//  new 了一个gin 的实例，没有加载handler
	r := gin.Default()

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.HttpAddr(),
		//Handler:           cors.AllowAll().Handler(r),
		Handler: r,
	}
	return &HttpService{
		server: server,
		l:      zap.L().Named("HTTP service"),
		r: r,
	}
}

type HttpService struct {
	server *http.Server
	l      logger.Logger
	r      gin.IRouter
}

func (s *HttpService) Start() error{

	//加载hanlder, 把handler注册给所有模块的gin
	apps.InitGin(s.r)
	//已加载app 的日志信息
	fmt.Println("开始注册1")
	apps := apps.LoadedGinApps()
	s.l.Infof("loaded gin apps:%v", apps)
	//该操作是一个阻塞的，简单端口， 等待请求
	//若为服务的正常关闭,
	fmt.Println("开始注册2")
	if err  := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed{
			s.l.Info("service stopped")
			return nil
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil

}
func (s *HttpService) Stop() error {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Errorf("graceful shutdown timeout, force exit",err)
	}
	return nil
}






















