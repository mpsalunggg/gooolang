package goweb

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "upload.gohtml", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	target, err := os.Create("resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	defer target.Close()

	if _, err := io.Copy(target, file); err != nil {
		panic(err)
	}

	name := r.FormValue("name")

	myTemplates.ExecuteTemplate(w, "upload_success.gohtml", map[string]interface{}{
		"Name":     name,
		"FileName": fileHeader.Filename,
		"FileUrl":  "/static/" + fileHeader.Filename,
	})
}

func TestUpload(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	writer.WriteField("name", "mps")

	part, _ := writer.CreateFormFile("file", "hello.txt")
	part.Write([]byte("hello upload file!"))

	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	result, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(result))
}

func TestServerUpload(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("resources"))))

	fmt.Println("open http://localhost:8181/")
	http.ListenAndServe("localhost:8181", mux)
}
