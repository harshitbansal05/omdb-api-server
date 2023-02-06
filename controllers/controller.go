package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/harshitbansal05/omdb-api-server/helpers"
	"github.com/joho/godotenv"
)

func init() {
	// Loading env
	cwd, _ := os.Getwd()
	err := godotenv.Load(cwd + "/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
}

func SearchMovies(c *gin.Context) {
	s := c.DefaultQuery("s", "")
	movieType := c.DefaultQuery("type", "")
	year := c.DefaultQuery("year", "")
	page := c.DefaultQuery("page", "1")

	resp, err := http.Get(fmt.Sprintf("%s?apikey=%s&s=%s&type=%s&y=%s&page=%s", os.Getenv("URL"), os.Getenv("ACCESS_KEY"), s, movieType, year, page))
	if err != nil {
		helpers.GenerateResponse(c, helpers.GETMOVIESFAILED, http.StatusBadRequest)
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		helpers.GenerateResponse(c, helpers.GETMOVIESFAILED, http.StatusBadRequest)
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(resBody)
}

func GetMovieDetailById(c *gin.Context) {
	id := c.Param("id")
	movieType := c.DefaultQuery("type", "")
	year := c.DefaultQuery("year", "")
	plot := c.DefaultQuery("plot", "short")

	resp, err := http.Get(fmt.Sprintf("%s?apikey=%s&i=%s&type=%s&y=%s&plot=%s", os.Getenv("URL"), os.Getenv("ACCESS_KEY"), id, movieType, year, plot))
	if err != nil {
		helpers.GenerateResponse(c, helpers.GETMOVIEDETAILFAILED, http.StatusBadRequest)
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		helpers.GenerateResponse(c, helpers.GETMOVIESFAILED, http.StatusBadRequest)
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(resBody)
}
