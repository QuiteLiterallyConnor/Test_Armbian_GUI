package main

import (
    "github.com/gin-gonic/gin"
    "os/exec"
    "runtime"
)

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.File("public/index.html")
    })

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
}

