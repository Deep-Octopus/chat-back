package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"go-chat/config"
	"go-chat/router"
	"go-chat/utils"
)

var (
	// specified config file path
	configFile string
)

func init() {
	//flag.StringVar(&configFile, "config", path.Join(filebox.GetCurrentRunningPath(), "config.yaml"), "app configuration file")
}
func main() {
	flag.Parse()
	//ctx := context.Background()
	configFile = "config\\config.yaml"
	appConfig := config.NewConfigFile(configFile)
	cfg := utils.InitConfig(appConfig)
	// set app mode
	gin.SetMode(cfg.App.Mode)

	utils.InitMySQL(cfg.Mysql)
	utils.InitRedis(cfg.Redis)
	r := router.Router()
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
