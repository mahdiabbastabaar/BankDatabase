package api

import (
	"BankDatabase/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var router *gin.Engine

func showAllSafeBoxes(context *gin.Context) {

	result, err := model.ShowAllSafeBoxes()
	if err.Err != nil {
		context.JSON(err.StatusCode, gin.H{
			"error": err.Err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, result)
	}
}

func showAllContracts(context *gin.Context) {
	result, err := model.ShowAllContracts()
	if err.Err != nil {
		context.JSON(err.StatusCode, gin.H{
			"error": err.Err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, result)
	}
}

func addSafeBox(context *gin.Context) {
	var safeBoxCreator model.SafeBox

	eror := context.ShouldBindJSON(&safeBoxCreator)
	if eror != nil {
		log.Println(eror)
	}

	safeBoxId, err := model.AddSafeBox(safeBoxCreator.MaximumValue, safeBoxCreator.CUId,
		safeBoxCreator.SHId, safeBoxCreator.PriceClass)
	if err.Err != nil {
		context.JSON(err.StatusCode, gin.H{
			"error": err.Err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"id": safeBoxId,
		})
	}

}

func editSafeBox(context *gin.Context) {
	var safeBoxCreator model.SafeBox

	eror := context.ShouldBindJSON(&safeBoxCreator)
	if eror != nil {
		log.Println(eror)
	}

	safeBoxId, err := model.EditSafeBox(int(safeBoxCreator.ID), safeBoxCreator.MaximumValue, safeBoxCreator.CUId,
		safeBoxCreator.SHId, safeBoxCreator.PriceClass)
	if err.Err != nil {
		context.JSON(err.StatusCode, gin.H{
			"error": err.Err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"id": safeBoxId,
		})
	}

}

func deleteSafeBox(context *gin.Context) {
	id := context.Param("id")
	safeBoxId, _ := strconv.Atoi(id)

	err := model.DeleteSafeBox(safeBoxId)
	if err.Err != nil {
		context.JSON(err.StatusCode, gin.H{
			"error": err.Err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, nil)
	}
}

func assignSafeBox(context *gin.Context) {
	type Ids struct {
		SafeBoxId  int `json:"safe_box_id"`
		CustomerId int `json:"customer_id"`
	}

	var ids Ids

	eror := context.ShouldBindJSON(&ids)
	if eror != nil {
		log.Println(eror)
	}

	id, err := model.AssignSafeBox(ids.SafeBoxId, ids.CustomerId)
	if err.Err != nil {
		context.JSON(err.StatusCode, gin.H{
			"error": err.Err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, id)
	}

}

func evacuateSafeBox(context *gin.Context) {
	type SafeBoxId struct {
		Id int `json:"id"`
	}

	var SBId SafeBoxId

	err := context.ShouldBindJSON(&SBId)
	if err != nil {
		log.Println(err)
	}

	model.EvacuateSafeBox()
}

func StartServer() {

	router = gin.Default()

	router.GET("/showsafeboxes", showAllSafeBoxes)
	router.GET("/currentcontracts", showAllContracts)
	router.POST("/addsafebox", addSafeBox)
	router.PATCH("/editsafebox", editSafeBox)
	router.DELETE("/deletesafebox/:id", deleteSafeBox)
	router.POST("/assignsafebox", assignSafeBox)
	router.POST("/evacuatesafeboxes", evacuateSafeBox)

	err := router.Run(":8000")
	if err != nil {
		return
	}
}
