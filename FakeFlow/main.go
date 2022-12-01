package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartTrainJob(id string) {
	exec.Command("/bin/bash", "file/pipelines/"+id+"/container.sh").Output()

}

func Unzip(fullpath string) {
	dst := "file/pipelines"
	archive, err := zip.OpenReader(fullpath)
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)
		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
			return
		}
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/pipeline/upload/:id", func(c *gin.Context) {
		file, _ := c.FormFile("file") // get file from form input name 'file'
		id := c.Param("id")
		c.SaveUploadedFile(file, "file/pipelines/"+id+".zip") // save file to tmp folder in current directory
		Unzip("file/pipelines/" + id + ".zip")
		os.Remove("file/pipelines/" + id + ".zip")
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	router.POST("pipeline/start/:id", func(c *gin.Context) {
		id := c.Param("id")
		StartTrainJob(id)
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	router.Run(":3500")
}
