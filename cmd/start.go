package cmd

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"

	"github.com/jayson-hu/api-demo-go/apps"
	_ "github.com/jayson-hu/api-demo-go/apps/all"
	"github.com/jayson-hu/api-demo-go/apps/host/http"
	"github.com/jayson-hu/api-demo-go/conf"
)

var (
	// pusher service config option
	confType string
	confFile string
	confETCD string
)

// StartCmd 程序的启动时 组装都在这里进行
// 1.
// StartCmd represents the base command when called without any subcommands
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 demo 后端API",
	Long:  "启动 demo 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载程序配置
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			panic(err)
		}
		//加载文件配置实体类

		//service := impl.NewHosServiceImpl()
		//注册hostservive 的实例注册到IOC
		//上面—— 导入
		//apps.HostService = impl.NewHosServiceImpl()
		//采用了 _ "github.com/jayson-hu/api-demo-go/apps/host/impl"

		//初始化配置日志模块初始化
		err = loadGlobalLogger()
		if err != nil {
			return err
		}

		//如何执行 hostservice 的config的方法 ， 没有包含func (i *HostServiceImpl) Config() {}的方法
		apps.InitImpl()
		//通过host API handler 提供 http restful
		//api := http.NewHostHTTPHander(service)
		api := http.NewHostHTTPHander()
		api.Config()
		//提供gin 的一个router
		g := gin.Default()
		//api.Register(g)
		apps.InitGin(g)
		g.Run(conf.C().App.HttpAddr())

		return errors.New("no flag find")
	},
}

// 问题：
// 1. http的 API GRPC api, 需要启动， 消息总线也需要监听，比如负责注册配置，这些模块都是独立
// 都需要再启动的时候，进行启动，都写在start， 会让start不易维护
// 2. 服务优雅的关闭？ 外部发动一个 terminal 中断信号
// 需要实现程序优雅关闭的逻辑的处理：由先后顺序(由外到内完成资源的是发放逻辑处理)
//		1. 包含api 层的关闭(http, grpc)
// 		2.消息总线的关闭
//		3.关闭数据库连接
//		3.如果使用了注册中心，最后才注销改实例，最后完成下线
//		4. 最后才退出完毕
// log 为全局变量, 只需要load 即可全局可用户, 依赖全局配置先初始化
func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)
	lc := conf.C().Log
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}
	zapConfig := zap.DefaultConfig()
	zapConfig.Level = level
	// 程序每启动一次，不必生成一个新的日志文件
	zapConfig.Files.RotateOnStartup = false
	// 配置日志的输出方式
	switch lc.To {
	case conf.ToStdout:
		zapConfig.ToStderr = true
		zapConfig.ToFiles = false
	case conf.ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}
	switch lc.Format {
	case conf.JSONFormat:
		zapConfig.JSON = true
	}
	// 把配置应用到全局logger
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}
	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}
func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}























