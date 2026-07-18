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

func TestCachedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*paths", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		paths := p.ByName("paths")

		fmt.Fprint(w, "Images :" + paths)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8181/images/user/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Images :/user/profile.png", string(body))
}
