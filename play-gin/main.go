package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	VER = "v0.0.13"
	dst = "./upgrade_pkg"
)

func main() {
	g := gin.Default()
	g.GET("/version", version)
	g.POST("/home", home)
	g.POST("/upload", upload)
	g.POST("/update", update)
	g.GET("/img", image)

	srv := &http.Server{
		Addr:    "192.168.2.106:9001",
		Handler: g,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

type Info struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func version(c *gin.Context) {
	c.String(http.StatusOK, "Version: %s", VER)
}

func home(c *gin.Context) {
	body := Info{}
	c.Bind(&body)
	fmt.Printf("%#v\n", body)
	c.String(http.StatusOK, "Maybe %v", "What's the format purpose?")
}

func update(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("form file error:", err)
	}

	destFile := dst + "/" + file.Filename
	err = c.SaveUploadedFile(file, destFile)
	if err != nil {
		log.Println("save file error: ", err)
	}

	switch runtime.GOOS {
	case "linux":
		linuxReload(destFile)
	case "windows":
		windowsReload(destFile)
	}

	c.String(http.StatusOK, "upgrade ok!")
}

func linuxReload(pkg string) {
	e, _ := os.Executable()
	os.Remove(e)

	unzip := exec.Command("unzip", pkg, "-d", ".")
	unzip.Run()

	kill := exec.Command("kill", "-2", fmt.Sprintf("%d", os.Getpid()))
	kill.Start()
}

func windowsReload(pkg string) {
	executable, _ := os.Executable()
	os.Rename(executable, "_deprecated.exe")

	unzip := exec.Command("unzip", pkg, "-d", ".")
	unzip.Run()

	os.Remove("_deprecated.exe")

	// sc := exec.Command("net", "stop", "play-gin")
	// sc.Start()
	// sc = exec.Command("sc", "start", "play-gin")
	// sc.Start()
}

func upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("form file error:", err)
	}

	os.Mkdir(dst, os.ModeDir)
	destFile := dst + "/" + file.Filename
	err = c.SaveUploadedFile(file, destFile)
	if err != nil {
		log.Println("save file error: ", err)
	}

	c.String(http.StatusOK, "upgrade ok!")
}

func image(c *gin.Context) {
	fileName := c.Query("img")
	if len(fileName) == 0 {
		log.Fatal("no filename")
	}

	c.File(dst + "/" + fileName)
}
