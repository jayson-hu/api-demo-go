package cmd

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jayson-hu/api-demo-go/apps"
	_ "github.com/jayson-hu/api-demo-go/apps/all"
	"github.com/jayson-hu/api-demo-go/apps/host/http"
	"github.com/jayson-hu/api-demo-go/conf"
	"github.com/spf13/cobra"
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

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
