package utils

import (
	"fmt"
	"math/rand"
	
	"time"

	
	"github.com/skip2/go-qrcode"
)

func GenerateQr(url string)(string,error){
	
	//Generate random file names
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	randomNumber := rng.Intn(1000)

	//Make a filepath 
	filename := fmt.Sprintf("qr_%d_%d.png",time.Now().UnixNano(), randomNumber)
	filepath := "images/" + filename

	err := qrcode.WriteFile(url, qrcode.Medium, 256, filepath)

	if err != nil {
		return "", err
	}

	return filename, nil
}