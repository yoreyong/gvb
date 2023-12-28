package router

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

// InitRouter Initialize gin router
func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	// Routers
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}

	// Settings api
	routerGroupApp.SettingsRouter()

	// ImagesUpload api
	routerGroupApp.ImagesRouter()

	return router
}
