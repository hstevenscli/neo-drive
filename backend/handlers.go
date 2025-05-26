package main



import (
	"github.com/gin-gonic/gin"
	"fmt"
)

var validPasswords = map[string]bool{
	"monke": true,
}

func postLogin(c *gin.Context) {
	var pw Password
	if err := c.BindJSON(&pw); err != nil {
		c.JSON(500, gin.H{"status": "Server error"})
		return
	}
	fmt.Println(pw)
	if validPasswords[pw.Password] == true {
		c.JSON(200, gin.H{"status": "success"})
		return
	}
	c.JSON(401, gin.H{"status": "incorrect password"})
}

