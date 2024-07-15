package controllers

import (
	"fmt"
	"items/models"

	"github.com/gin-gonic/gin"
)

type (
	controllers struct {
		repository models.MysqlDatabase
	}

	Controllers interface {
		GetItems(ctx *gin.Context)
		CreateItems(ctx *gin.Context)
	}
)

func InitControllers(db models.MysqlDatabase) Controllers {
	fmt.Println("<<< Init Controller >>>")
	return &controllers{
		repository: db,
	}
}
