package main

import "github.com/webview/webview"

func main() {
    debug := true
    w := webview.New(debug)
    defer w.Destroy()
    w.SetTitle("Example WebView Application")
    w.SetSize(800, 600, webview.HintNone)
    // Load a local HTML file or use w.Navigate("http://example.com") to load a URL
    w.Navigate("./public/index.html")
    w.Run()
}
