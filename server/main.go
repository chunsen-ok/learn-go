package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type LoginInfo struct {
	Name     string `json:"name"`
	Password string `json:"pass"`
}

func main() {
	srv := gin.Default()
	srv.StaticFile("/", "web/index.html")

	srv.POST("/login", login)

	srv.Run(":8544")
}

func login(c *gin.Context) {
	m := LoginInfo{}
	if err := c.ShouldBindJSON(&m); err != nil {
		log.Println("Login error:", err)
	}

	fmt.Println("login:", m)
}
