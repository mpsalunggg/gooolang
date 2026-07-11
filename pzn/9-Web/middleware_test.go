package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

// ServeHTTP membuat LogMiddleware memenuhi interface http.Handler.
func (l *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before executing handler:", r.URL.Path)

	l.Handler.ServeHTTP(w, r) // teruskan ke handler asli

	fmt.Println("After executing handler:", r.URL.Path)
}

func MiddlewareHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Middleware")
}

func TestMiddleware(t *testing.T) {
	handler := &LogMiddleware{
		Handler: http.HandlerFunc(MiddlewareHandler),
	}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181/", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println("Response body:", string(body))
}

func TestServerMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", MiddlewareHandler)

	handler := &LogMiddleware{
		Handler: mux,
	}

	http.ListenAndServe("localhost:8181", handler)
}

func TestServerMiddleware2(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler execute")
		fmt.Fprintf(w, "Hello gaiss")
	})

	logMiddleware := new(LogMiddleware)
	logMiddleware.Handler = mux

	server := http.Server{
		Addr:    "localhost:8181",
		Handler: logMiddleware,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

type ErrorHandler struct {
	Handler http.Handler
}

func (e *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v", err)
		}
	}()

	e.Handler.ServeHTTP(w, r)
}

func PanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("erorr nich")
}

func TestErrorHandler(t *testing.T) {
	handler := &ErrorHandler{
		Handler: http.HandlerFunc(PanicHandler),
	}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181/", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println("Status : ", response.StatusCode)
	fmt.Println("Body   :", string(body))
}

func NormalHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func TestErrorHandlerNormal(t *testing.T) {
	handler := &ErrorHandler{
		Handler: http.HandlerFunc(NormalHandler),
	}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181/", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println("Status :", response.StatusCode)
	fmt.Println("Body   :", string(body))
}

func TestServerErrorHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic", PanicHandler)
	mux.HandleFunc("/oke", NormalHandler)

	var handler http.Handler = mux
	handler = &LogMiddleware{Handler: handler}
	handler = &ErrorHandler{Handler: handler}

	http.ListenAndServe("localhost:8181", handler)
}
