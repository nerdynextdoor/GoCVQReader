package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"time"

	"github.com/liyue201/goqr"
)

func recognizeFile(path string) {
	start := time.Now()
	fmt.Printf("recognize file: %v\n", path)
	imgdata, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	img, _, err := image.Decode(bytes.NewReader(imgdata))
	if err != nil {
		fmt.Printf("image.Decode error: %v\n", err)
		return
	}
	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		fmt.Printf("Recognize failed: %v\n", err)
		return
	}

	for _, qrCode := range qrCodes {

		fmt.Printf("qrCode text: %s\n", qrCode.Payload)

	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func main() {
	recognizeFile("QRgrid.png")
}
