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

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/upload", upload)
	r.HandleFunc("/download", download)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

	fmt.Println("exit")
}
