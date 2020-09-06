package main

import (
	"fmt"
	_ "log"
	"net/http"
	"github.com/gin-gonic/gin"
)

const dst = "files"

func main() {
	g := gin.Default()
	g.POST("/home", home)
	g.Run("127.0.0.1:9001")
}


type Info struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}

// /home
func home(c *gin.Context) {
	body := Info {}
	c.Bind(&body)
	fmt.Printf("%#v\n", body)
	c.String(http.StatusOK, "Maybe %v", "What's the format purpose?")
}
