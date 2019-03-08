package main

import (
	"log"
	"syscall/js"
)

// main requires Go 1.12+.
func main() {

	doc := js.Global().Get("document")
	log.Println("WASM main works")

	clickBtn := js.FuncOf(func(this js.Value, evt []js.Value) interface{} {
		log.Println("clicked")
		return nil
	})

	btn := doc.Call("getElementById", "try-button")
	btn.Call("addEventListener", "click", clickBtn)

	select {}
}
