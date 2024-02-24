package main

import rl "github.com/uprun/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
        //var runes = []rune{'q','w','e','r','t','y','u','i','o','p','[',']','{','}','й','ц','у','к','е','н','г','ш','щ','з','х','ї','ф','і','в','а','п','р','о','л','д','ж','є','ʼ','ґ','я','ч','с','м','и','т','ь','б','ю','.'}
        var font = rl.LoadFontExByRunesNumber("./hack-font/Hack-Regular.ttf",32, 1256)


	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
                var position = rl.Vector2 {
                    X: 190,
                    Y: 200,
                }
		rl.DrawTextEx(font, "Congrats! You cre(ated your Тепер перевіремо українську!", position, 20.0, 3.0, rl.LightGray)

		rl.EndDrawing()
	}
}