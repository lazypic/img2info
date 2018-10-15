package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
	"log"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("blender_opengl_test.png")
	text, err := client.Text()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)
}
