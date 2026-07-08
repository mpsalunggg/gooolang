package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("function.gohtml").ParseFiles("./templates/function.gohtml"))

	t.ExecuteTemplate(w, "function.gohtml", map[string]any{
		"Title": "Template Function",
		"SayHello": func(name string) string {
			return "Hello, " + name + "!"
		},
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("function_global.gohtml").ParseFiles("./templates/function_global.gohtml"))

	t.ExecuteTemplate(w, "function_global.gohtml", map[string]any{
		"Title": "Template Function Global",
		"Name":  "Mps",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateFunctionCustom(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("function_custom.gohtml").Funcs(template.FuncMap{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	}).ParseFiles("./templates/function_custom.gohtml"))

	t.ExecuteTemplate(w, "function_custom.gohtml", map[string]any{
		"Title": "Template Function Custom",
		"Name":  "Mps",
	})
}

func TestTemplateFunctionCustom(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCustom(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
