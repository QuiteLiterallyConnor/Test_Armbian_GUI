package main

import (
    "os/exec"
)

func main() {
    htmlFilePath := "./public/index.html"
    
    cmd := exec.Command("chromium-browser", "--kiosk", htmlFilePath)
    err := cmd.Start()
    if err != nil {
        panic(err)
    }
}
