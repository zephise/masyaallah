package models

import (
	"fmt"
	cModels "items/controllers/models"
	"items/models/mapping"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	mysqlDatabase struct {
		DB *gorm.DB
	}

	MysqlDatabase interface {
		CreateItems(ctx *gin.Context, data mapping.Items) error
		GetItems(ctx *gin.Context, params cModels.ParamsGetItems) ([]mapping.Items, int64, error)
	}
)

func InitDatabase() MysqlDatabase {
	fmt.Println("<<<< Initialize Database Connection >>>>")
	return &mysqlDatabase{
		DB: ConnectionMysql(),
	}
}

var logMode = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"error":  logger.Error,
	"warn":   logger.Warn,
	"info":   logger.Info,
}

func ConnectionMysql() *gorm.DB {
	// 	DATABASE_USERNAME=root
	// DATABASE_PASSWORD=Zxcv0987654321@
	// DATABASE_HOST=localhost
	// DATABASE_PORT=3306
	// DATABASE_NAME=item
	// DEBUG_MYSQL="true"
	// LOG_MODE_MYSQL="info"
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	debug := os.Getenv("DATABASE_DEBUG_MYSQL")
	mode := os.Getenv("LOG_MODE_MYSQL")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logMode[mode]),
	})

	if err != nil {
		fmt.Println("Error Connect Database", err)
		panic("Error Connection Database")
	}

	fmt.Println("Mysql Connected Successfully")

	if debug == "true" {
		return db.Debug()
	}
	return db
}
