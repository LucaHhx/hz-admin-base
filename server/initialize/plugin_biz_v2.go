package initialize

import (
	plugin "hab/utils/plugin/v2"

	"github.com/gin-gonic/gin"
)

func PluginInitV2(group *gin.Engine, plugins ...plugin.Plugin) {
	for i := 0; i < len(plugins); i++ {
		plugins[i].Register(group)
	}
}

func bizPluginV2(engine *gin.Engine) {
}
