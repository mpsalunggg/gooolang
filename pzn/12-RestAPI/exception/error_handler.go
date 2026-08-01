package exception

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/model/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}

	if validationErrors(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func writeError(w http.ResponseWriter, status int, statusText string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	webResponse := web.WebResponse{
		Code:   status,
		Status: statusText,
		Data:   data,
	}

	if err := json.NewEncoder(w).Encode(webResponse); err != nil {
		fmt.Printf("failed to encode error response: %v\n", err)
	}
}

func validationErrors(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if !ok {
		return false
	}

	writeError(w, http.StatusBadRequest, "BAD REQUEST", exception.Error())
	return true
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if !ok {
		return false
	}

	writeError(w, http.StatusNotFound, "NOT FOUND", exception.Error)
	return true
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	writeError(w, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", fmt.Sprintf("%v", err))
}
