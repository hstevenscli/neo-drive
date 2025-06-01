package main



import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"os"
	"log"
)

type Password struct {
	Password string `json:"password"`
}

type file struct {
	Filename string `json:"filename"`
}

type fileObj struct {
	Files []string `json:"files"`
}

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


func handleFileUpload(c *gin.Context) {
	form, err := c.MultipartForm()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid upload"})
        return
    }

    files := form.File["files"] // must match the name in formData.append("files", ...)
    for _, file := range files {
        // Save each file or process as needed
        err := c.SaveUploadedFile(file, "./uploads/" + file.Filename)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
            return
        }
    }
    c.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfully"})
}


func getFiles(c *gin.Context) {
	files, err := os.ReadDir("./uploads/")
	var rfiles []string
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		rfiles = append(rfiles, file.Name())
	}
	fmt.Println(rfiles)

	// files := fileObj{ []string{ "file1", "file2" } }
	c.JSON(200, fileObj{ rfiles })
}

func getFileByName(c *gin.Context) {
    filename := c.Param("filename")
    c.Header("Content-Disposition", "attachment; filename=" + filename)
    c.File("./uploads/" + filename)
}

func deleteFileByName(c *gin.Context) {
	filename := c.Param("filename")
	err := os.Remove("./uploads/"+filename)
	if err != nil {
		c.JSON(400, gin.H{"error": "couldn't delete file"})
	}
	c.JSON(201, gin.H{"status": "Successfully deleted file: "+ filename})
}
