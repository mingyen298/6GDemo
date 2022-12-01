package main

import (
	"net/http"
	"os/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTrainJob(id string) {
	exec.Command("/bin/bash", "./run.sh", id).Output()
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("train_job/create", func(c *gin.Context) {

		uuid := uuid.NewString()
		CreateTrainJob(uuid)

		c.JSON(http.StatusOK, gin.H{"message": "success", "uuid": uuid})
	})

	router.Run(":3499")
}
