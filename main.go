package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func GenerateQr(qr string) (string, error) {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	randomNumber := rng.Intn(1000)

	filename := fmt.Sprintf("qr_%d_%d.png", time.Now().UnixNano(), randomNumber)
	filepath := "images/" + filename

	err := qrcode.WriteFile(qr, qrcode.Medium, 256, filepath)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func QrHandler(c *gin.Context) {
	url := c.DefaultQuery("url", "")
	fmt.Println("Test:", url)

	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "enter valid url"})
		return
	}

	filename, err := GenerateQr(url)
	fmt.Println("Generated filename:", filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot generate qr"})
		return
	}

	imageURL := fmt.Sprintf("http://localhost:9090/images/%s", filename)
	c.JSON(http.StatusOK, gin.H{"message": "QR code generated successfully", "qr_image_url": imageURL})
}

func main() {

	router := gin.Default()

	router.Static("/images", "./images")
	router.GET("/generate", QrHandler)
	router.Run(":9090")
}
