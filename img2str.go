package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
	"github.com/tuotoo/qrcode"
	"log"
	"os"
)

func main() {
	// text
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("blender_opengl_test.png")
	text, err := client.Text()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)
	//
	fi, err := os.Open("blender_opengl_test.png")
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
