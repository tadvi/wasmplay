package main

import (
	"log"
	"syscall/js"

	"github.com/tadvi/plate"
)

var doc = js.Global().Get("document")

type App struct {
	Enter   int
	KeyDown string

	Values []string

	Checkbox bool
}

var app App

// main requires Go 1.12+.
func main() {
	log.Println("WASM main")

	/*clickBtn := js.FuncOf(func(this js.Value, evt []js.Value) interface{} {
		log.Println("clicked")
		return nil
	})

	btn := doc.Call("getElementById", "try-button")
	btn.Call("addEventListener", "click", clickBtn)*/

	p := plate.New("my")
	app.Enter = 11
	app.KeyDown = "works"
	app.Values = []string{"1", "2", "3", "4"}

	p.Render(&app)

	select {}
}
