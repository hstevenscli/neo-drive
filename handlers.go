package main

import (
	"mime"
	"fmt"
	"strings"
	"net/http"
	"os"
	"syscall"
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
	".js": true,
}

func postLogin(c *gin.Context) {
	var pw Password
	if err := c.BindJSON(&pw); err != nil {
		c.JSON(500, gin.H{"status": "Server error"})
		return
	}
	password := os.Getenv("PASSWORD")
	fmt.Println("got pw", pw)
	fmt.Println("password:", password)
	if pw.Password == password {
		c.JSON(200, gin.H{"status": "success"})
		return
	}
	c.JSON(401, gin.H{"status": "incorrect password"})
}


func handleFileUpload(c *gin.Context) {
	path := c.Param("path")
	form, err := c.MultipartForm()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid upload"})
        return
    }

    files := form.File["files"] // must match the name in formData.append("files", ...)
    for _, file := range files {
        // Save each file or process as needed
        err := c.SaveUploadedFile(file, "./uploads/" + path + "/" + file.Filename)
		fmt.Println("Path:", path)
		fmt.Println("Filename:", file.Filename)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
            return
        }
    }
    c.JSON(201, gin.H{"message": "Files uploaded successfully"})
}


//TODO in progress changing filenames []string -> entries []dirEntry
func readDirectory(c *gin.Context, path string) {
	fmt.Println("Getting directory")
	files, err := os.ReadDir("./uploads/" + path)
	fmt.Println("files:", files)
	var filenames []string
	var entries []dirEntry
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
		de := dirEntry{
			Name: file.Name(),
			IsDir: file.IsDir(),
		}
		entries = append(entries, de)
		filenames = append(filenames, file.Name())
	}
	fmt.Println(filenames)
	// files := fileObj{ []string{ "file1", "file2" } }
	c.JSON(200, entries)
}

// func getFileByName(c *gin.Context) {
//     filename := c.Param("filename")
//     c.Header("Content-Disposition", "attachment; filename=" + filename)
//     c.File("./uploads/" + filename)
// }

func getFileByName(c *gin.Context) {
    filename := c.Param("filename")
    path := "./uploads/" + filename

    // Ensure EPUB is sent correctly
    if strings.HasSuffix(strings.ToLower(filename), ".epub") {
        c.Header("Content-Type", "application/epub+zip")
    }

    c.Header(
        "Content-Disposition",
        fmt.Sprintf(`attachment; filename="%s"`, filename),
    )

    c.File(path)
}

// takes in the full path to the current file, and a new name to change it in place
// (doenst change location)
func renameFile(c *gin.Context) {
	path := c.Param("path")
	fmt.Println("path:", path)
	type Request struct {
		Newname string `json:"newname"`
	}
	r := Request{}
	c.BindJSON(&r)
	fmt.Println("newname:", r)
	// func Rename(oldpath, newpath string) error
	pathSlice := strings.Split(path, "/")
	fmt.Println("slice:", pathSlice)
	newPath := strings.Join(pathSlice[:len(pathSlice) - 1], "/")
	newPathComplete := fmt.Sprintf("%s/%s", newPath, r.Newname)
	fmt.Println("newPath:", newPathComplete)

	err := os.Rename("./uploads/" + path, "./uploads" + newPathComplete)


	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(204, gin.H{"status": "successfully renamed file"})
}

func postNewDir(c *gin.Context) {
	path := c.Param("path")
	err := os.Mkdir("./uploads/" + path, 0755)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "Error creating directory", 
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	c.JSON(201, gin.H{"status": "Created directory"})
	fmt.Println("MKDIR PATH", path)
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
	err := os.Remove("./uploads/" + filename)
	if err != nil {
		fmt.Println(err)
		if pathErr, ok := err.(*os.PathError); ok {
			fmt.Println(pathErr.Err)
			if pathErr.Err == syscall.ENOTEMPTY {
				fmt.Println("NOTEMPTRY")
				c.JSON(400, gin.H{"error": "Directory not empty"})
				return
			}
		}
		c.JSON(400, gin.H{"error": "couldn't delete file"})
		return
	}
	c.JSON(200, gin.H{"status": "Successfully deleted file: "+ filename})
	return
}
