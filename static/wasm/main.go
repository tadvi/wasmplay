package main

import (
	"fmt"
	"log"
	"syscall/js"

	"github.com/tadvi/browser/ajax"
	"github.com/tadvi/browser/dom"
)

var doc = js.Global().Get("document")

type App struct {
}

var app App

// main requires Go 1.12+.
func main() {
	log.Println("WASM main")

	dom.Select("#submit").On("click", func(this, evt js.Value) {

		fmt.Println("clicked")
		ajax.Post("/ajax", "", func(resp string) {
			fmt.Println(resp)
		})
	})

	select {}
}
