package main

import rl "github.com/gen2brain/raylib-go/raylib"
import (
//   "io/ioutil"
    "log"
    "net/http"
    "golang.org/x/net/html"
    "fmt"
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
    resp, err := http.Get("https://news.ycombinator.com/")
    defer resp.Body.Close()
    
    if err != nil {
      log.Fatalln(err)
    }
    doc, err := html.Parse(resp.Body)
    if err != nil {
      log.Fatalln(err)
    }

    RenderNode(doc, 0)
}

func RenderNode(node *html.Node, level int) {
    //fmt.Println("type:", node.Type)
    for lvl := 0; lvl < level; lvl ++ {
        fmt.Print("-")
    }
    fmt.Println("Data:", node.Data)
    // traverse children
    for c := node.FirstChild; c != nil; c = c.NextSibling {
        RenderNode(c, level + 1)
    }
}