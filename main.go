package main

import (
	"container/list"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	// "io"
	// "bytes"
	// "strings"
)

type node struct {
	name	string
	list	list.List
	// next	*node
}

func readFileChunkWise(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 1048576)  // обработка 1048576 байтов файла
	file1, header, err := r.FormFile("file")
	file1.Close()
	file, err := os.Open(header.Filename);
	if err != nil {
		fmt.Printf("File not found\n")
		return
	}

	list := list.New()
	// var fileSave node = node{header.Filename, *list}
	os.Mkdir("createdDir", 0777)
	i := 1
	for {
		bytesRead, _ := file.Read(b);
		if bytesRead == 0 { // bytesRead будет равен 0 в конце файла.
			break
		}
		name := strconv.Itoa(i) + "-" + header.Filename
		dirName := "createdDir/" + name
		process(b, bytesRead, dirName); // создание файла с порядковым именем
		list.PushBack(name)
		// fmt.Printf("%s\n",list)
		i++
	}
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)  // печать списка файлов по 1 мб
	}
	file.Close();
	postform(w, r)
}

func process(b []byte, bytesRead int, dirName string) {
	err := ioutil.WriteFile(dirName, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}



func postform(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("sources_html/postform.html")
	t.Execute(w, "")
}

func download(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("sources_html/download.html")
	// file1, header, _ := r.FormFile("file")
	// file1.Close()
	// t.ExecuteTemplate(w, "", header.Filename)
	t.Execute(w, "")
}

func handler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("sources_html/upload.html")
	t.Execute(w, "")

	fmt.Println("Метод:", r.Method)
	if r.Method == "POST" {
		readFileChunkWise(w, r)
	} else if r.Method == "GET" {
		fmt.Println("go in get")

	} else {
		fmt.Println("** error **")
	}
}

func main() {
	http.HandleFunc("/upload", handler)
	http.HandleFunc("/download", download)
	http.HandleFunc("/postform", postform)
	http.ListenAndServe(":8080", nil)

	fmt.Println("exit")
}
