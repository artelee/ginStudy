package main

import (
	api "ginStudy/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "머글")

		c.String(http.StatusOK, api.GetWelcomeMessage(name))
	})
	router.Run(":8080")
}
