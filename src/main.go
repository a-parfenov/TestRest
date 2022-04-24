package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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

	r := mux.NewRouter()
	r.HandleFunc("/upload", handler.upload)
	r.HandleFunc("/download", handler.download)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

	fmt.Println("exit")
}
