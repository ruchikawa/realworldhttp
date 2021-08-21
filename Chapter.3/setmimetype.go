package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)


func main() {

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("Name","Michael Jackson")

	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/png")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.png"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}

	readFile, err := os.Open("photo.png")
	if err != nil {
		panic(err)
	}
	
	defer readFile.Close()
	io.Copy(fileWriter,readFile)
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
