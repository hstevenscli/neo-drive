package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
    "time"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

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
        MaxAge: 12 * time.Hour,
    }))

	r.Static("/assets", "./frontend/dist/assets/")
	// GET ROUTES
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})
	r.GET("/dir", func(c *gin.Context) {
		readDirectory(c, "")
	})
	r.GET("/dir/*path", func(c *gin.Context) {
		path := c.Param("path")
		readDirectory(c, path)
	})


	r.GET("/files/*filename", getFileByName)
	r.GET("/view/*filename", viewFile)


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
	r.POST("/upload/*path", handleFileUpload)
	r.POST("/dir/*path", postNewDir)

	// DELETE ROUTES
	r.DELETE("files/*filename", deleteFileByName)

    r.Run("0.0.0.0:"+port)
}
