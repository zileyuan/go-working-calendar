package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zileyuan/go-working-calendar/controller"
	"github.com/zileyuan/go-working-calendar/util"
)

// AuthTokenRoutes 路由的AuthToken状态表
var AuthTokenRoutes = make(map[string]bool)

// Route 路由管理
func Route(g *gin.Engine) {
	g.GET("/holiday/:date", controller.HolidayAction)
	g.GET("/workingdays/:from/:to", controller.CountAction)
	g.GET("/workingdate/:from/:amount", controller.CalcAction)

	for k, v := range AuthTokenRoutes {
		util.Infof("Route: %v; Auth: %v\n", k, v)
	}
}
