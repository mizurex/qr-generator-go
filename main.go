package main

import (
	"qr-generator/handlers"
	"github.com/gin-gonic/gin"
	
)

func main() {

	router := gin.Default()

	router.Static("/images", "./images")
	router.GET("/generate",handlers.QrHandler)
	router.Run(":9090")
}
