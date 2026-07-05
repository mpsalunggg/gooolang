package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Address struct {
	Street string
}

type Data struct {
	Title   string
	Name    string
	Address Address
}

func DataTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/data.gohtml"))

	t.ExecuteTemplate(w, "data.gohtml", map[string]interface{}{
		"Title": "Helloo",
		"Name":  "Mps",
		"Address": Address{
			Street: "Belum ada",
		},
	})
}

func TestDataTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181", nil)
	recorder := httptest.NewRecorder()

	DataTemplate(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
