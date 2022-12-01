package main

import (
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func DeploymentModel(id string) {
	exec.Command("/bin/bash", `deployment/run.sh`, id).Output()
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/model/upload/:id", func(c *gin.Context) {
		file, _ := c.FormFile("file") // get file from form input name 'file'
		id := c.Param("id")
		c.SaveUploadedFile(file, "models/"+id+".zip") // save file to tmp folder in current directory

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	router.GET("/model/download/:id", func(c *gin.Context) {
		id := c.Param("id")
		fileType := ".zip"
		_, errByOpenFile := os.Open("models/" + id + fileType)

		if errByOpenFile != nil {
			c.Redirect(http.StatusFound, "/404")
			return
		}
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+id+fileType)
		c.Header("Content-Transfer-Encoding", "binary")
		c.File("models" + "/" + id + fileType)
		// return
	})

	router.POST("/model/deployment/:id", func(c *gin.Context) {
		id := c.Param("id")
		DeploymentModel(id)

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	router.Run(":3501")
}
