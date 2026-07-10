package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")

	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "BAD REQUEST")
		return
	}

	path := "./resources/" + filename

	if _, err := os.Stat(path); err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "FILE NOT FOUND: %s", filename)
		return
	}

	w.Header().Add("Content-Disposition", "attachment; filename=\""+filename+"\"")
	http.ServeFile(w, r, path)
}

func TestDownloadFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181/download?file=text.txt", nil)
	recorder := httptest.NewRecorder()

	DownloadFile(recorder, request)

	response := recorder.Result()
	fmt.Println("Status              :", response.StatusCode)
	fmt.Println("Content-Disposition :", response.Header.Get("Content-Disposition"))

	body, _ := io.ReadAll(response.Body)
	fmt.Println("Body                :", string(body))
}

func TestServerDownload(t *testing.T) {
	http.HandleFunc("/download", DownloadFile)
	fmt.Println("open http://localhost:8181/download?file=text.txt")
	http.ListenAndServe("localhost:8181", nil)
}
