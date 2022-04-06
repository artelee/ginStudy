package router

import (
	"fmt"
	"ginStudy/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCommonRouter(port int) {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "머글")

		c.String(http.StatusOK, api.GetWelcomeMessage(name))
	})
	router.Run(`:` + strconv.Itoa(port))
	fmt.Printf(`---port:%d로 서버실행---`, port)
}
