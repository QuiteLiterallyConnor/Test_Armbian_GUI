package main

import (
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve specific route for index.html
	r.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	// Serve other static files
	r.Static("/public", "./public")

	port := "8080"
	url := "http://localhost:" + port

	go func() {
		if err := r.Run(":" + port); err != nil {
			panic(err)
		}
	}()

	runtime.Gosched()

	cmd := exec.Command("chromium-browser", "--no-sandbox", "--kiosk", url)
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	println("Shutting down server...")
}