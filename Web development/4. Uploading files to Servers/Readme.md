
<!-- https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types -->
<!-- example of MIME type aka Media type: Content-Type", "application/json;charset='utf-8" -->


<!-- https://freshman.tech/file-upload-golang/#report-the-upload-progress -->
1. Just like http.Request we have http.Reponse struct which is popluated by http.ResponseWriter interface 
2. http.maxBytesReader only caps the size of body.
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
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

## ParseForm function https://pkg.go.dev/net/http#Request.ParseForm
---------------------------------------
- ParseForm populates r.Form and r.PostForm.
- For all requests, ParseForm parses the raw query from the URL and updates r.Form.
- For POST, PUT, and PATCH requests, it also reads the request body, parses it as a form and puts the     results into both r.PostForm and r.Form. Request body parameters take precedence over URL query string values in r.Form.
- If the request Body's size has not already been limited by MaxBytesReader, the size is capped at 10MB.
- For other HTTP methods, or when the Content-Type is not application/x-www-form-urlencoded, the request Body is not read, and r.PostForm is initialized to a non-nil, empty value.
- ParseMultipartForm calls ParseForm automatically. ParseForm is idempotent.



// A response represents the server side of an HTTP response.
type response struct {
	conn             *conn
	req              *Request // request for this response
	reqBody          io.ReadCloser
	cancelCtx        context.CancelFunc // when ServeHTTP exits
	wroteHeader      bool               // reply header has been (logically) written
	wroteContinue    bool               // 100 Continue response was written
	wants10KeepAlive bool               // HTTP/1.0 w/ Connection "keep-alive"
	wantsClose       bool               // HTTP request has Connection "close"

	w  *bufio.Writer // buffers output in chunks to chunkWriter
	cw chunkWriter

	// handlerHeader is the Header that Handlers get access to,
	// which may be retained and mutated even after WriteHeader.
	// handlerHeader is copied into cw.header at WriteHeader
	// time, and privately mutated thereafter.
	handlerHeader Header
	calledHeader  bool // handler accessed handlerHeader via Header

	written       int64 // number of bytes written in body
	contentLength int64 // explicitly-declared Content-Length; or -1
	status        int   // status code passed to WriteHeader

	// close connection after this reply.  set on request and
	// updated after response from handler if there's a
	// "Connection: keep-alive" response header and a
	// Content-Length.
	closeAfterReply bool

	// requestBodyLimitHit is set by requestTooLarge when
	// maxBytesReader hits its max size. It is checked in
	// WriteHeader, to make sure we don't consume the
	// remaining request body to try to advance to the next HTTP
	// request. Instead, when this is set, we stop reading
	// subsequent requests on this connection and stop reading
	// input from it.
	requestBodyLimitHit bool

	// trailers are the headers to be sent after the handler
	// finishes writing the body. This field is initialized from
	// the Trailer response header when the response header is
	// written.
	trailers []string

	handlerDone atomicBool // set true when the handler exits

	// Buffers for Date, Content-Length, and status code
	dateBuf   [len(TimeFormat)]byte
	clenBuf   [10]byte
	statusBuf [3]byte

	// closeNotifyCh is the channel returned by CloseNotify.
	// TODO(bradfitz): this is currently (for Go 1.8) always
	// non-nil. Make this lazily-created again as it used to be?
	closeNotifyCh  chan bool
	didCloseNotify int32 // atomic (only 0->1 winner should send)
}