package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
    "time"
	"fmt"

)


func main() {
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

	// GET ROUTES
	r.GET("/files", getFiles)
	r.GET("/files/:filename", getFileByName)
	r.GET("/view/:filename", viewFile)

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
	r.POST("/upload", handleFileUpload)

	// DELETE ROUTES
	r.DELETE("files/:filename", deleteFileByName)

	r.Run("localhost:8080")
}
