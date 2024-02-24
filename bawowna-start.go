package main

import rl "github.com/uprun/raylib-go/raylib"
import (
   "io/ioutil"
   "log"
   "net/http"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
        var font = rl.LoadFontExByRunesNumber("./hack-font/Hack-Regular.ttf",32, 1256)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
                var position = rl.Vector2 {
                    X: 190,
                    Y: 200,
                }
		rl.DrawTextEx(font, "Congrats! You created", position, 20.0, 3.0, rl.LightGray)
                position = rl.Vector2 {
                    X: 190,
                    Y: 240,
                }
		rl.DrawTextEx(font, "Тепер перевіремо українську!", position, 20.0, 3.0, rl.LightGray)

		rl.EndDrawing()
	}
}

func HttpGetHackerNews() {
    // this is test code I copied
    resp, err := http.Get("https://news.ycombinator.com/")
    if err != nil {
      log.Fatalln(err)
    }
        //We Read the response body on the line below.
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
              log.Fatalln(err)
        }
        //Convert the body to type string
        sb := string(body)
        log.Printf(sb)
        // end of test code
}