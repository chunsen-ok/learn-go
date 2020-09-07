package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/gin-gonic/gin"
)

const dst = "./upgrade_pkg"

func main() {
	g := gin.Default()
	g.GET("/version", version)
	g.POST("/home", home)
	g.POST("/upload", upload)
	g.Run("127.0.0.1:9001")
}

type Info struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func version(c *gin.Context) {
	c.String(http.StatusOK, "Version: V0.0.1")
}

func home(c *gin.Context) {
	body := Info{}
	c.Bind(&body)
	fmt.Printf("%#v\n", body)
	c.String(http.StatusOK, "Maybe %v", "What's the format purpose?")
}

func upload(c *gin.Context) {
	h := gin.H{}

	file, err := c.FormFile("file")
	if err != nil {
		log.Println("form file error:", err)
	}

	destFile := dst + "/" + file.Filename
	err = c.SaveUploadedFile(file, destFile)
	if err != nil {
		log.Println("save file error: ", err)
	}
	h["upload_file"] = file.Filename

	fmt.Println("go os:", runtime.GOOS)
	switch runtime.GOOS {
	case "windows":
		executable, _ := os.Executable()
		os.Rename(executable, "other.exe")
	case "linux":
		// ...
	}

	unzip := exec.Command("unzip", destFile, "-d", "./")
	unzip.Run()

	c.JSON(http.StatusOK, h)

	switch runtime.GOOS {
	case "windows":
		os.Remove("other.exe")

		sc := exec.Command("net", "stop", "play-gin")
		sc.Start()
		// sc = exec.Command("sc", "start", "play-gin")
		// sc.Start()

	case "linux":
		systemctl := exec.Command("systemtcl", "restart", "sycc")
		systemctl.Start()
	}
}
