package helpers

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/harshitbansal05/omdb-api-server/models"
)

func GenerateResponse(c *gin.Context, message string, status int) {
	resp := models.ResponseJSONModel{Message: message}
	data, err := json.Marshal(&resp)
	if err != nil {
		log.Print(err)
		return
	}
	c.Writer.WriteHeader(status)
	c.Writer.Write(data)
}
