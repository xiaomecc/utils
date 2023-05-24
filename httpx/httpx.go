package httpx

import (
	"bytes"
	"net/http"
)

type ResponseRecorder struct {
	W          http.ResponseWriter
	StatusCode int
	Body       *bytes.Buffer
}

func NewRecorder(w http.ResponseWriter) *ResponseRecorder {
	return &ResponseRecorder{
		W:          w,
		StatusCode: http.StatusOK,
		Body:       new(bytes.Buffer),
	}
}

func (rw *ResponseRecorder) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
	rw.W.WriteHeader(statusCode)
}

func (rw *ResponseRecorder) Write(buf []byte) (int, error) {
	_, _ = rw.Body.Write(buf)
	return rw.W.Write(buf)
}

func (rw *ResponseRecorder) Header() http.Header {
	return rw.W.Header()
}
