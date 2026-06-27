package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	w.Header().Add("X-Powered-By", "mps")
	fmt.Fprint(w, contentType)
}

func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
	fmt.Println(recorder.Header().Get("x-powered-by"))

}
