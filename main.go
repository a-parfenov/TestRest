package main

import (
	"bytes"
	"container/list"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	// "strings"
	// "bytes"
	// "strings"
)

type node struct {
	name	string
	save	map[string] bool
	// list	list.List
	// next	*node
}

func readFileChunkWise(w http.ResponseWriter, r *http.Request) {
	var size int = 300
	b := make([]byte, size)  // обработка 1048576 байтов файла
	// for i := 0; i < size; i++ {
	// 	b[i] = 0
	// }
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
		for i := 0; i < size; i++ {
			b[i] = 0
		}
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
	// strings.TrimRight()
	// strings.TrimPrefix()
	// bytes.NewBuffer()
	
	file.Close();
	io.WriteString(w, `
		<form >
			<h3>Файл сохранен</h3>
		</form>
		`)
}

func process(b []byte, bytesRead int, dirName string) {
	err := ioutil.WriteFile(dirName, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func existFinc(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true }
	if os.IsNotExist(err) {
		return false }
	return false
}

func postform(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("sources_html/postform.html")
	t.Execute(w, "")
}

func download(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("sources_html/download.html")
	t.Execute(w, "")

	// if existFinc("createdDir"){
	// 	fmt.Fprintf(w, "Сборка и сохранение файла")

		// _, header, _ := r.FormFile("text")
		// fmt.Fprintf(w, header.Filename)
		// _, err := os.Open("createdDir/1-" + header.Filename);
		// if err != nil {
		// 	fmt.Fprintf(w, "Такой файл не сохранен")
		// 	return
		// } else {
			// b := make([]byte, 1048576)
			// for i := 0; i < 1000; i++ {
			// 	file, err := os.Open("createdDir/" + strconv.Itoa(i) + "-" + header.Filename);
			// 	if err != nil { break }

			// 	i := 1
			// 	for {
			// 		bytesRead, _ := file.Read(b);
			// 		if bytesRead == 0 { break }
			// 		dirName := header.Filename
			// 		process(b, bytesRead, dirName);
			// 		i++
			// 	}
			// }
		}

	// } else {
	// 	io.WriteString(w, `
	// 	<form >
	// 		<h3>Файл не сохранен</h3>
	// 	</form>
	// 	`)
	// }

	// file1, header, _ := r.FormFile("file")
	// file1.Close()
	// t.ExecuteTemplate(w, "", header.Filename)
// }

func handler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("sources_html/upload.html")
	t.Execute(w, "")

	fmt.Println("Метод:", r.Method)
	if r.Method == "POST" {
		// fmt.Fprintf(w, "go in post")
		readFileChunkWise(w, r)
	} else if r.Method == "GET" {
		// fmt.Fprintf(w, "go in get")
		// download(w, r)

	   	// crutime := time.Now().Unix()
	   	// h := md5.New()
	   	// io.WriteString(h, strconv.FormatInt(crutime, 10))
	   	// token := fmt.Sprintf("%x", h.Sum(nil))

	   	// t, _ := template.ParseFiles("upload.gtpl")
	   	// t.Execute(w, token)
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
