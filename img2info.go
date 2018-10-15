package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
	"github.com/tuotoo/qrcode"
	"log"
	"os"
)

func main() {
	testfile := "blender_opengl_test.png"
	// text
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(testfile)
	text, err := client.Text()
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(text)
	// qrcode
	fi, err := os.Open(testfile)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(qrmatrix.Content)
}
