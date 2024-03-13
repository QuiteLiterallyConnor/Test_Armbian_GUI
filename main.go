package main

import (
    "github.com/gin-gonic/gin"
    "os/exec"
    "os"
    "os/signal"
    "syscall"
    "runtime"
)

func main() {
    // Initialize the Gin engine.
    r := gin.Default()

    // Serve the HTML file.
    r.GET("/", func(c *gin.Context) {
        c.File("./index.html") // Ensure you have an index.html file in the same directory.
    })

    // Determine the port to listen on. Ensure this matches the URL you open in Chromium.
    port := "8080"
    url := "http://localhost:" + port

    // Run the server in a goroutine so it doesn't block the subsequent execution.
    go func() {
        if err := r.Run(":" + port); err != nil {
            panic(err)
        }
    }()

    // Give the server a moment to start.
    runtime.Gosched()

    // Open Chromium to the URL.
    cmd := exec.Command("chromium-browser", "--no-sandbox", "--kiosk", url)
    if err := cmd.Start(); err != nil {
        panic(err)
    }

    // Wait for interrupt signal to gracefully shut down the server.
    quit := make(chan os.Signal, 1)
    // Catch both SIGINT (Ctrl+C) and SIGTERM (termination signal from systemd).
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    // Add any cleanup logic here. For example, you might want to close database connections,
    // flush logs, or other graceful shutdown tasks.
    println("Shutting down server...")
}
