package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB
)

// http.Response
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method passed by browser: ", r.Method)
	var bs []byte
	// acting as Gaurd clause
	if r.Method != http.MethodPost {
		http.Error(w, "Method must be POST", http.StatusInternalServerError)
		return
	}
	// NOTE: http.maxBytesReader only caps the size of body.
	// requestTooLarge is called by maxBytesReader when too much input has been read from the client.
	// func (w *response) requestTooLarge() {
	// 	w.closeAfterReply = true
	// 	w.requestBodyLimitHit = true // ATTENTION
	// 	if !w.wroteHeader {
	// 		w.Header().Set("Connection", "close")
	// 	}
	// }
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	//  If this method returns an error, then the size limit was breached or the request body was invalid in some
	if err := r.ParseForm(); err != nil {
		fmt.Println("Error parsing file.")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	file, fileheader, err := r.FormFile("file") // very important.
	if err != nil {
		log.Printf("Error uploading file. %s", err.Error())
	}
	defer file.Close()
	fmt.Println("File header: ", fileheader)
	bs, _ = ioutil.ReadAll(file)
	fmt.Println("File Content: ", string(bs))

	w.Header().Set("Content-type", "text/html")
	io.WriteString(w, `
		<form method="POST" enctype="multipart/form-data">
 	 	<input type="file" name="arun">
  		<input type="submit">
		</form>
		<br>`+string(bs))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func main() {
	mx := mux.NewRouter()
	mx.HandleFunc("/", indexHandler)
	mx.HandleFunc("/upload", uploadHandler)
	log.Fatal(http.ListenAndServe(":9000", mx))
}
