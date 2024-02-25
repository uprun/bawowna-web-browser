package main

import rl "github.com/gen2brain/raylib-go/raylib"
import (
    "log"
    "net/http"
    "golang.org/x/net/html"
    "fmt"
    "io"
    "bytes"
)

func main() {
	rl.InitWindow(800, 450, "Bawowna web browser")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
        var font = rl.LoadFontEx("./hack-font/Hack-Regular.ttf",32, nil, 1256)
        defer rl.UnloadFont(font)
        var position = rl.Vector2 {
            X: 190,
            Y: 200,
        }

        HttpGetHackerNews()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
                mouseDiff := rl.GetMouseWheelMoveV()
                position.X += mouseDiff.X
                position.Y += mouseDiff.Y
                
		rl.DrawTextEx(font, "Congrats! You created", position, 20.0, 3.0, rl.LightGray)
                position2 := rl.Vector2 {
                    X: position.X,
                    Y: position.Y + 30,
                }
		rl.DrawTextEx(font, "Тепер перевіремо українську!", position2, 20.0, 3.0, rl.LightGray)

		rl.EndDrawing()
	}
}

func HttpGetHackerNews() {
    client := &http.Client{}
    req, _ := http.NewRequest("GET", "https://news.ycombinator.com/", nil)
    req.Header.Set("User-Agent", "Bawowna-web-browser?v=2024-02-25")
    resp, err := client.Do(req)
    defer resp.Body.Close()
    
    if err != nil {
      log.Fatalln(err)
    }

// Read the response body
    body_bytes, _ := io.ReadAll(resp.Body)
 
    // Print the body as a string
    fmt.Println("HTML:\n\n", string(body_bytes))

    reader := bytes.NewReader(body_bytes)
    doc, err := html.Parse(reader)
    if err != nil {
      log.Fatalln(err)
    }

    RenderNode(doc, 0)
    PrettyRenderNode(doc)
}

func PrettyRenderNode(node *html.Node) {
    //fmt.Println("type:", node.Type)
    if node.Type == html.ElementNode && node.Data == "tr" {
        fmt.Println("")
    }

    if node.Type == html.ElementNode && node.Data == "br" {
        fmt.Println("")
    }

    if node.Type == html.ElementNode && node.Data == "div" {
        fmt.Println("")
    }

    if node.Type == html.ElementNode && node.Data == "td" {
        fmt.Print("|")
    }

    if node.Type != html.ElementNode  {
        fmt.Print(node.Data)
    }
    
    
    // traverse children
    for c := node.FirstChild; c != nil; c = c.NextSibling {
        PrettyRenderNode(c)
    }
}

func RenderNode(node *html.Node, level int) {
    //fmt.Println("type:", node.Type)
    for lvl := 0; lvl < level; lvl ++ {
        fmt.Print("-")
    }
    fmt.Println("Data:", node.Data)
    for _, attribute := range node.Attr {
        for lvl := 0; lvl <= level; lvl ++ {
            fmt.Print("-")
        }
        fmt.Printf("%q = %q\n", attribute.Key, attribute.Val)
    }
    // traverse children
    for c := node.FirstChild; c != nil; c = c.NextSibling {
        RenderNode(c, level + 1)
    }
}