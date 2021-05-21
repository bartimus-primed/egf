package main

import (
	"encoding/json"
	"fmt"
	"log"

	layout "github.com/bartimus-primed/egf/core/types"
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	ss := layout.Create_Flexbox()

	screenWidth := int32(1024)
	screenHeight := int32(768)

	rl.SetConfigFlags(rl.FlagVsyncHint)

	rl.InitWindow(screenWidth, screenHeight, "raylib [gui] example - basic controls")
	var inputText1 string
	var inputText2 string
	var inputText3 string
	var inputText4 string

	x1, y1, width1, height1 := ss.Add_Item(0, 0, 200, 20)
	x2, y2, width2, height2 := ss.Add_Item(0, 0, 200, 20)
	x3, y3, width3, height3 := ss.Add_Item(0, 0, 200, 20)
	x4, y4, width4, height4 := ss.Add_Item(0, 0, 200, 20)
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.Beige)
		inputText1 = rg.TextBox(rl.NewRectangle(x1, y1, width1, height1), inputText1)
		inputText2 = rg.TextBox(rl.NewRectangle(x2, y2, width2, height2), inputText2)
		inputText3 = rg.TextBox(rl.NewRectangle(x3, y3, width3, height3), inputText3)
		inputText4 = rg.TextBox(rl.NewRectangle(x4, y4, width4, height4), inputText4)

		rl.EndDrawing()
	}
	ssJSON, err := json.MarshalIndent(ss, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%s", string(ssJSON))
	rl.CloseWindow()
}
