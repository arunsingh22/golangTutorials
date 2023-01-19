package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	var bs []byte
	if r.Method == http.MethodPost {
		f, fh, err := r.FormFile("arun")
		if err != nil {
			log.Printf("Error uploading file.")
		}
		defer f.Close()
		fmt.Println("File header: ", fh)
		bs, _ = ioutil.ReadAll(f)
		fmt.Println("File Content: ", string(bs))
	}

	w.Header().Set("Content-type", "text/html;charset=utf-8")
	io.WriteString(w, `
		<form method="POST" enctype="multipart/form-data">
 	 	<input type="file" name="arun">
  		<input type="submit">
		</form>
		<br>`+string(bs))
}

func main() {
	http.HandleFunc("/", upload)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
