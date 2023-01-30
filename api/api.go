package api

import "github.com/gin-gonic/gin"

var router *gin.Engine

func showAllSafeBoxes(context *gin.Context) {

}

func showAllContracts(context *gin.Context) {

}

func addSafeBox(context *gin.Context) {

}

func editSafeBox(context *gin.Context) {

}

func deleteSafeBox(context *gin.Context) {

}

func assignSafeBox(context *gin.Context) {

}

func evacuateSafeBox(context *gin.Context) {

}

func StartServer() {

	router = gin.Default()

	router.GET("/showsafeboxes", showAllSafeBoxes)
	router.GET("/currentcontracts", showAllContracts)
	router.GET("/addsafebox", addSafeBox)
	router.PATCH("/editsafebox", editSafeBox)
	router.DELETE("/deletesafebox", deleteSafeBox)
	router.GET("/assignsafebox", assignSafeBox)
	router.GET("/evacuatesafeboxes", evacuateSafeBox)

	err := router.Run(":8000")
	if err != nil {
		return
	}
}
