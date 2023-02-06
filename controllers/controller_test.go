package controllers_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/harshitbansal05/omdb-api-server/routes"
)

func TestSearchMoviesController(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search?s=hello", nil)
	router.ServeHTTP(w, req)

	resp, _ := http.Get(fmt.Sprintf("%s?apikey=%s&s=hello", os.Getenv("URL"), os.Getenv("ACCESS_KEY")))
	resBody, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 200, w.Code)
	wBody, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, resBody, wBody)
}

func TestIncorrectIDSearchMoviesController(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search?s=", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	wBody, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, []byte("{\"Response\":\"False\",\"Error\":\"Incorrect IMDb ID.\"}"), wBody)
}

func TestTooManyResultsSearchMoviesController(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search?s=s", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	wBody, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, []byte("{\"Response\":\"False\",\"Error\":\"Too many results.\"}"), wBody)
}

func TestGetMovieDetailController(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/detail/tt3766394", nil)
	router.ServeHTTP(w, req)

	resp, _ := http.Get(fmt.Sprintf("%s?apikey=%s&i=tt3766394", os.Getenv("URL"), os.Getenv("ACCESS_KEY")))
	resBody, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 200, w.Code)
	wBody, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, resBody, wBody)
}

func TestQueryParamsGetMovieDetailController(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/detail/tt3766394?plot=full", nil)
	router.ServeHTTP(w, req)

	resp, _ := http.Get(fmt.Sprintf("%s?apikey=%s&i=tt3766394&plot=full", os.Getenv("URL"), os.Getenv("ACCESS_KEY")))
	resBody, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 200, w.Code)
	wBody, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, resBody, wBody)
}

func TestIncorrectIdGetMovieDetailController(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/detail/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	wBody, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, []byte("{\"Response\":\"False\",\"Error\":\"Incorrect IMDb ID.\"}"), wBody)
}

func TestEmptyIdGetMovieDetailController(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/detail/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}
