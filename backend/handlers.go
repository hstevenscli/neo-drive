package main

import (
	"mime"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
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

type dirEntry struct {
	Name string
	IsDir bool
}

var validPasswords = map[string]bool{
	"monke": true,
}

var viewableFileTypes = map[string]bool {
	".txt": true,
	".html": true,
	".md": true,
	".cpp": true,
	".py": true,
	".cxx": true,
	".c": true,
	".h": true,
	".csv": true,
	".png": true,
	".jpg": true,
	".jpeg": true,
	".pdf": true,
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
	var rfiles []dirEntry
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

func viewFile(c *gin.Context) {
	filename := c.Param("filename")

	// verify valid viewable file
	ext := filepath.Ext(filename)
	fmt.Println("EXT:", ext)
	if !viewableFileTypes[ext] {
		c.JSON(415, gin.H{"error": "Unsupported file type for preview"})
		return
	}
	// if viewable get the file data and send it
	data, err := os.ReadFile("./uploads/" + filename)
	if err != nil {
		c.JSON(400, gin.H{"error": "File Not Found"})
		return
	}

	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = http.DetectContentType(data)
	}

	c.Header("Content-Type", mimeType)
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(data)
}

func deleteFileByName(c *gin.Context) {
	filename := c.Param("filename")
	err := os.Remove("./uploads/"+filename)
	if err != nil {
		c.JSON(400, gin.H{"error": "couldn't delete file"})
	}
	c.JSON(201, gin.H{"status": "Successfully deleted file: "+ filename})
}
