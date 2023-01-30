package api

import "github.com/gin-gonic/gin"

var router *gin.Engine

func StartServer() {

	router = gin.Default()

	router.GET()
	router.POST()

	err := router.Run(":8000")
	if err != nil {
		return
	}
}
