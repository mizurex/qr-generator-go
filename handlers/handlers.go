package handlers

import (
	"fmt"
	"net/http"
	
	"qr-generator/utils"
	"github.com/gin-gonic/gin"
	
)

func QrHandler(c *gin.Context){
	//Take out the url
	url := c.DefaultQuery("URL","")

	if url == ""{
		c.JSON(http.StatusBadRequest, gin.H{"message": "enter valid url"})
		return
	}

	filename, err := utils.GenerateQr(url)

	fmt.Println("Generated filename:", filename)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot generate qr"})
		return
	}

	imageUrl := fmt.Sprintf("http://localhost:9090/images/%s", filename)

	c.JSON(http.StatusOK, gin.H{"message": "QR code generated successfully", "qr_image_url": imageUrl})
	

}