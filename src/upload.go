package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func readFileChunk(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("%s", err)
		return
	}
	defer func() { file.Close() }()

	os.Mkdir(dirSaved, 0777)
	var i int = 0

	for {
		b := make([]byte, sizeBuf)

		_, err := file.Read(b)
		if err != nil {
			if err.Error() == "EOF" { break }
			log.Fatal(err)
			return
		}
		creatingAFile(b, i, header.Filename)  // Creating a file with an ordinal name
		i++
	}
	io.WriteString(w, `<form><h3>Файл сохранен</h3></form>`)
}

func (h *handler) upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/upload.html")
	t.Execute(w, "")

	fmt.Println("Метод:", r.Method)
	if r.Method == "POST" {
		readFileChunk(w, r)
	}
}

func creatingAFile(b []byte, i int, fileName string) {
	name := fileName + "-" + strconv.Itoa(i)

	err := ioutil.WriteFile(dirSaved + "/" + name, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
