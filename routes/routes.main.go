package routes

import (
	"fmt"
	ctrl "items/controllers"

	"github.com/gin-gonic/gin"
)

type (
	Router struct {
		Controllers ctrl.Controllers
		Gin         *gin.Engine
	}
	RouterInterface interface {
		StartServer()
	}
)

func InitRoutes(ctrl ctrl.Controllers) RouterInterface {
	return &Router{
		Controllers: ctrl,
		Gin:         gin.Default(),
	}
}

func (r *Router) StartServer() {
	fmt.Println("Initialize Router")

	items := r.Gin.Group("/items")
	{
		items.POST("/", r.Controllers.CreateItems)
		items.GET("/", r.Controllers.GetItems)
	}

	r.Gin.Run(":8000")
}
