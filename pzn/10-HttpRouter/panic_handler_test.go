package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error : ", err)
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Ups")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
	assert.Equal(t, "Error : Ups", string(body))
}
