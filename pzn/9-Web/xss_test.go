package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	err := myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title":   "Hellooo",
		"Content": "<p>welcommm</p>",
	})
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDisableEscape(w http.ResponseWriter, r *http.Request) {
	err := myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title":   "Hellooo",
		"Content": template.HTML("<p>welcommm</p>"),
	})
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

func TestTemplateDisableEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181", nil)
	recorder := httptest.NewRecorder()

	TemplateDisableEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateXSS(w http.ResponseWriter, r *http.Request) {
	userInput := r.URL.Query().Get("content")

	err := myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title":   "XSS Demo",
		"Content": template.HTML(userInput),
	})
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

func TestTemplateXSS(t *testing.T) {
	payload := "<script>alert('XSS')</script>"
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181?content="+payload, nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println("=== template.HTML (VULNERABLE) ===")
	fmt.Println(string(body))
}

func TestTemplateXSSProtected(t *testing.T) {
	payload := "<script>alert('XSS')</script>"
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181?content="+payload, nil)
	recorder := httptest.NewRecorder()

	userInput := request.URL.Query().Get("content")
	myTemplates.ExecuteTemplate(recorder, "post.gohtml", map[string]interface{}{
		"Title":   "XSS Protected",
		"Content": userInput,
	})

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println("=== plain string (SAFE) ===")
	fmt.Println(string(body))
}

func TestServerXSS(t *testing.T) {
	http.HandleFunc("/", TemplateXSS)
	fmt.Println("open http://localhost:8181/?content=<script>alert('XSS')</script>")
	http.ListenAndServe("localhost:8181", nil)
}
