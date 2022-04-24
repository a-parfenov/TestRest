package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func checkFiles(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true }
	if os.IsNotExist(err) {
		return false }
	return false
}

func checkPartFile(partFile []byte) ([]byte) {
	var size int = sizeBuf
	for i := 0; i < size; i++ {
		if partFile[i] == 0 {
			size = i
		}
	}

	if size == sizeBuf {
		return partFile
	} else {
		tmpPartFile := make([]byte, size)
		for j := 0; j < size; j++ {
			tmpPartFile[j] = partFile[j]
		}
		return tmpPartFile
	}
}

func handDownload(w http.ResponseWriter, r *http.Request, fileName string) {
	var fullFile bytes.Buffer
	var i int = 0

	for {
		partFile, err := os.ReadFile(dirSaved + "/" + fileName + "-" + strconv.Itoa(i))
		if err != nil {
			if strings.Contains(err.Error(), "no such file or directory") {
				break
			}
			log.Fatal(err)
		}
		partFile = checkPartFile(partFile)
		fullFile.Write(partFile)  // We collect all parts of the file in the buffer
		i++
	}

	if checkFiles(dirDownload + "/" + fileName){
		io.WriteString(w, `<form><h3>Ошибка: файл существует</h3></form>`)
		return
	}
	os.Mkdir(dirDownload, 0777)
	file, err := os.Create(dirDownload + "/" + fileName)
	if err != nil{
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(fullFile.String())
	io.WriteString(w, `<form><h3>Файл собран и сохранен</h3></form>`)
}

func (h *handler) download(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/download.html")
	t.Execute(w, "")

	fileName := r.FormValue("text")
	if len(fileName) == 0 { return }
	
	tmpFile, err := os.Open(dirSaved + "/" + string(fileName) + "-0")
	if err != nil {
		log.Print("file not found")
		io.WriteString(w, `<form><h3>Файл не найден</h3></form>`)
		w.WriteHeader(404)
		return
	}
	tmpFile.Close()
	handDownload(w, r, fileName)
}
