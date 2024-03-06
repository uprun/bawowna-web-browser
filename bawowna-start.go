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

type RenderRequest struct {
    Text string
    Anchor rl.Vector2 
}

var queue = make([]RenderRequest, 0)
var font rl.Font

func main() {
	rl.InitWindow(800, 450, "Bawowna web browser")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
        font = rl.LoadFontEx("./hack-font/Hack-Regular.ttf",32, nil, 1256)
        defer rl.UnloadFont(font)
        var position = rl.Vector2 {
            X: 0,
            Y: 0,
        }

        HttpGetHackerNews()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
                mouseDiff := rl.GetMouseWheelMoveV()
                position.X += mouseDiff.X * 2
                position.Y += mouseDiff.Y * 2
 
                for _, top_request := range queue {

                    top_request.Anchor.X += position.X
                    top_request.Anchor.Y += position.Y
                    var offset float32 = 0.0
                    // draw individual glyph/rune separately
                    for _, v := range top_request.Text {
                        rl.DrawTextEx(font, string([]rune {v}), rl.Vector2{top_request.Anchor.X + offset, top_request.Anchor.Y} , 20.0, 3.0, rl.Green)
                        offset += rl.MeasureTextEx(font, string([]rune {v}), 20.0, 3.0).X + 3.0
                    }
                }


		rl.EndDrawing()
	}
}

func HttpGetHackerNews() {
    client := &http.Client{}
    req, _ := http.NewRequest("GET", "https://news.ycombinator.com/", nil)
    req.Header.Set("User-Agent", "Bawowna-web-browser?v=2024-03-03")
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

    TreeRenderNode(doc, 0)
    GuiRenderNode(doc, rl.Vector2{0,0})
}


func GuiRenderNode(node *html.Node, anchor_parent rl.Vector2) rl.Vector2 {
    anchor := anchor_parent
    if node.Type == html.ElementNode && node.Data == "tr" {
        anchor.X = 0
        anchor.Y += 20
    }

    if node.Type == html.ElementNode && node.Data == "br" {
        anchor.X = 0
        anchor.Y += 20
    }

    if node.Type == html.ElementNode && node.Data == "div" {
        anchor.X = 0
        anchor.Y += 20
    }

    if node.Type == html.ElementNode && node.Data == "td" {
        anchor.X += 20
    }
//    fmt.Println(anchor_parent, anchor)

    if node.Type != html.ElementNode  {
        queue = append(queue, RenderRequest{node.Data, anchor})
        size := rl.MeasureTextEx(font, node.Data, 20.0, 3.0)
        anchor.X += size.X + 20
    }
    
    
    // traverse children
    for c := node.FirstChild; c != nil; c = c.NextSibling {
        anchor = GuiRenderNode(c, anchor)
    }
    return anchor
}

func TreeRenderNode(node *html.Node, level int) {
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
        TreeRenderNode(c, level + 1)
    }
}