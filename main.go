package main

import (
    "fmt"
    "os/exec"
)

func main() {
    htmlFilePath := "./public/index.html"
    
    cmd := exec.Command("chromium-browser", "--kiosk", htmlFilePath)
    
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Error running command: %v\n", err)
    }
    
    fmt.Printf("Command output: %s\n", output)
}