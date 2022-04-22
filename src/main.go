package main

import (
	"fmt"
	"net/http"
)

const (
	dirSaved	= "savedFiles"
	dirDownload	= "downloadFiles"
	sizeBuf		= 1048576  // 1 mb == 1048576 bytes
)

type handler struct {
	savesFiles map[string]bool
}

func NewHandler() *handler {
	return &handler{
		savesFiles: make(map[string]bool),
	}
}

func main() {
	handler := NewHandler()

	http.HandleFunc("/upload", handler.upload)
	http.HandleFunc("/download", handler.download)

	http.ListenAndServe(":8080", nil)

	fmt.Println("exit")
}
