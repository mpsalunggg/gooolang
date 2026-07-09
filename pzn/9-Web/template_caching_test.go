package goweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//go:embed templates/*.gohtml
var templatesCaching embed.FS

var myTemplates = template.Must(template.New("templates").Funcs(template.FuncMap{
	"upper": func(value string) string {
		return strings.ToUpper(value)
	},
	"sayHello": func(value string) string {
		return "Hello, " + value + "!"
	},
}).ParseFS(templatesCaching, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.gohtml", "Hello Template Caching")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
