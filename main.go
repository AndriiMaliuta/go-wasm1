package main

import (
	"flag"
	"fmt"
	"syscall/js"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
)

func add(i []js.Value) {
	js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
}

func registerCallbacks() {
	//js.Global().Set("add", js.NewCallback(add))
}

func getElem(name string) {
	root := js.Global().Get("document").Call("querySelector", name).Get("value")
	println(root)
}

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("Hello, WebAssembly!")
	flag.Parse()
	registerCallbacks()

	getElem("root")

	<-c
}
