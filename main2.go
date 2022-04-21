package main


	import (

	  "net/http"
	  "io"
	  "os"
	  "path/filepath"
	  "fmt"
	)
	func main1() {

	  fmt.Println("TEMP DIR:", os.TempDir())
	  http.ListenAndServe(":9000", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
		  src, hdr, err := req.FormFile("my-file")
		  if err != nil {
			http.Error(res, err.Error(), 500)
			return
		  }
		  defer src.Close()

		  dst, err := os.Create(filepath.Join(os.TempDir(), hdr.Filename))
		  if err != nil {
			http.Error(res, err.Error(), 500)
			return
		  }
		  defer dst.Close()

		  io.Copy(dst, src)
		}

		res.Header().Set("Content-Type", "text/html")
		io.WriteString(res, `
		  <form method="POST" enctype="multipart/form-data">
			<input type="file" name="my-file">
			<input type="submit">
		  </form>
		  `)
		  file, err := os.Open("1.txt")
		  if err != nil {
			  return
		  }
		  defer file.Close()
	  
		  stat, err := file.Stat()
		  if err != nil {
			  return
		  }
		  bs := make([]byte, stat.Size())
		  _, err = file.Read(bs)
		  if err != nil {
			  return
		  }
	  
		  str := string(bs)
		  fmt.Println(str)

		}))
		

		// http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request){
		// 	fmt.Fprint(w, "Contact Page")
		// })

		// type page struct {
		// 	title	string
		// 	body	[]byte
		// }


		// func () HandlerSave(file) {
		// 	save()
		// }

		// func (p *page) save() os.Error {
		// 	filename := p.title + ".txt"
		// 	return ioutil.WriteFile(filename, p.body, 0600)
		// }
	}
	
