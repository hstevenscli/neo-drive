package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	// load env variable
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, continuing with environment variables")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	// Configure CORS options
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Static("/assets", "./frontend/dist/assets/")
	// GET ROUTES
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})
	r.GET("/dir", authorizeMiddleware(), func(c *gin.Context) {
		readDirectory(c, "")
	})
	r.GET("/dir/*path", authorizeMiddleware(), func(c *gin.Context) {
		path := c.Param("path")
		readDirectory(c, path)
	})
	r.GET("/files/*filename", authorizeMiddleware(), getFileByName)
	r.GET("/view/*filename", authorizeMiddleware(), viewFile)
	r.GET("/session", getSession)

	// POST ROUTES
	r.POST("/ping", func(c *gin.Context) {
		var pw Password
		if err := c.BindJSON(&pw); err != nil {
			c.JSON(401, gin.H{"status": "got error"})
			return
		}
		fmt.Println(pw)
		c.JSON(200, gin.H{"status": "got this password:" + pw.Password})
	})
	r.POST("/login", postLogin)
	r.POST("/upload/*path", authorizeMiddleware(), handleFileUpload)
	r.POST("/dir/*path", authorizeMiddleware(), postNewDir)

	r.PUT("/rename/*path", authorizeMiddleware(), renameFile)

	// DELETE ROUTES
	r.DELETE("files/*filename", authorizeMiddleware(), deleteFileByName)

	r.Run("0.0.0.0:" + port)
}
