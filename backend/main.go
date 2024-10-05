package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	shellout "shellout/utils"
)

var db *gorm.DB

func setupDB() {
	var err error

	var (
		POSTGRES_HOST     = os.Getenv("POSTGRES_HOST")
		POSTGRES_PORT     = os.Getenv("POSTGRES_PORT")
		POSTGRES_USER     = os.Getenv("POSTGRES_USER")
		POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
		POSTGRES_DB       = os.Getenv("POSTGRES_DB")
	)

	dsn := "host=" + POSTGRES_HOST + " user=" + POSTGRES_USER + " password=" + POSTGRES_PASSWORD + " dbname=" + POSTGRES_DB + " port=" + POSTGRES_PORT + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Database connected!")
}

func main() {
	router := gin.Default()

	setupDB()

	// Diagnostic
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// API

	// Templates
	router.GET("api/templates", shellout.GetTemplateList)
	// router.GET("/templates/:id", shellout.GetTemplate)
	// router.POST("/templates", shellout.CreateTemplate)
	// router.PUT("/templates/:id", shellout.UpdateTemplate)
	// router.DELETE("/templates/:id", shellout.DeleteTemplate)

	// // Terms
	// router.GET("/terms", shellout.GetTermList)
	// router.GET("/terms/:id", shellout.GetTerm)
	// router.POST("/terms", shellout.CreateTerm)
	// router.PUT("/terms/:id", shellout.UpdateTerm)
	// router.DELETE("/terms/:id", shellout.DeleteTerm)

	// // Docker Hosts
	// router.GET("/dockerhosts", shellout.GetDockerHostList)
	// router.POST("/dockerhosts", shellout.CreateDockerHost)
	// router.GET("/dockerhosts/:id", shellout.ConnectDockerHost)
	// router.DELETE("/dockerhosts/:id", shellout.DisconnectDockerHost)
	// router.DELETE("/dockerhosts/:id", shellout.DeleteDockerHost)

	// // Instances
	// router.GET("/instances", shellout.GetInstanceList)
	// router.GET("/instances/:id", shellout.GetInstance)
	// router.POST("/instances", shellout.CreateInstance)
	// router.PUT("/instances/:id", shellout.UpdateInstance)
	// router.DELETE("/instances/:id", shellout.DeleteInstance)
	// router.POST("/instances/:id/start", shellout.StartInstance)
	// router.POST("/instances/:id/stop", shellout.StopInstance)
	// router.POST("/instances/:id/restart", shellout.RestartInstance)

	router.Run(":8080")
}
