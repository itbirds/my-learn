package main

import (
	"io"
	"log"
	"net/http"
)

func hellos(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w,"hello world222")
}

func main() {
	http.HandleFunc("/09",hellos)
	err := http.ListenAndServe(":8080",nil)
	if err!=nil {
		log.Fatal("listenandserver",err.Error())
	}
}

