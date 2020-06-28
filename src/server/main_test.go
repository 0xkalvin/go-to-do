package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"go-to-do/src/database"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
 }

 func TestListTodo(t *testing.T) {

	var database = database.InitializeDatabase()
	var router = SetupRouter(database)

	w := performRequest(router, "GET", "/todos")

	assert.Equal(t, http.StatusOK, w.Code,  "It should return status code 200")
 }