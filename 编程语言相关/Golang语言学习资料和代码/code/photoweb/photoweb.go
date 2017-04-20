package main

import(
	"io"
	"log"
	"net/http"
)

func uploadHandler(w http.ResponseWriter,r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w,"<html><body><form method=\"POST\" action=\"/upload\" enctype=\"multipart/form-data\">" +
			"choose an image to upload：<input name=\"image\" type=\"file\" />" +
			"<input type=\"submit\" value=\"upload\" />" +
			"</form></body></html>")
		return
	}
}

func main() {
	http.HandleFunc("/upload",uploadHandler)
	err := http.ListenAndServe(":8080",nil)
	if err!=nil {
		log.Fatal("listenandserver",err.Error())
	}
}