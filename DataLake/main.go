package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file") // get file from form input name 'file'

		c.SaveUploadedFile(file, "tmp/"+file.Filename) // save file to tmp folder in current directory

		c.String(http.StatusOK, "file: %s", file.Filename)
	})

	router.Run()
}
