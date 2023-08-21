package controllers

import (
	"mini-challenge-pertemuan-sembilan/database"
	"mini-challenge-pertemuan-sembilan/helpers"
	models "mini-challenge-pertemuan-sembilan/models/entity"
	requests "mini-challenge-pertemuan-sembilan/models/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()

	var productReq requests.ProductRequest
	// Form, use ShouldBind
	if err := ctx.ShouldBind(&productReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract the filename without extension
	fileName := helpers.RemoveExtension(productReq.Image.Filename)

	uploadResult, err := helpers.UploadFile(productReq.Image, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Product := models.Product{
		Name:        productReq.Name,
		Description: productReq.Description,
		Stock:       productReq.Stock,
		ImageURL:    uploadResult,
	}

	err = db.Debug().Create(&Product).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Product,
	})
}
