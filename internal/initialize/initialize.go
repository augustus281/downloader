package initialize

import (
	"fmt"
	"github.com/augustus281/downloader/global"
)

func Initalize() {
	LoadConfig()
	InitLogger()
	r := InitRouter()
	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)
	if global.Config.Server.Mode != "release" {
		fmt.Println(serverAddr)
	}
	r.Listen(serverAddr)
}
