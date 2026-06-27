package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(w, "hello")
	} else {
		fmt.Fprintf(w, "hello %s", name)
	}
}

func TestQuery(testing *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181/hello?name=mps", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func SayMultipleHello(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "%s %s", firstName, lastName)
}

func TestMultipleQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181/hello?first_name=mps&last_name=al", nil)
	recorder := httptest.NewRecorder()

	SayMultipleHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)

}
