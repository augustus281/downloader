package initialize

import (
	"github.com/augustus281/downloader/global"
	"github.com/augustus281/downloader/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
