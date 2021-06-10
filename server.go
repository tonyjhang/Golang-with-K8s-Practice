package main
import (
	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
)

type RestServer struct {
	port string
}

func (s *RestServer) run() {
	r := gin.Default()
	r.GET("/", index)
	r.POST("/test", test)
	fmt.Println("Server has init.")
	r.Run(s.port)
}
func index(c *gin.Context) {
	c.JSON(200, "A Golang Rest Practice.")
}

func test(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		c.JSON(503, gin.H{
			"message": "there's internal error occurs.",
		})
	}
	fmt.Println("Get Request:", string(body))
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
