package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	var document = js.Global.Get("document")
	var body = document.Get("body")

	var three = js.Global.Get("THREE")
	//var scene = three.Get("Scene").New()
	//var camera = three.Get("PerspectiveCamera").New(75, 500, 0.1, 1000.0)
	var renderer = three.Get("WebGLRenderer").New()

	renderer.Call("setSize", 500, 500)

	body.Call("appendChild", renderer.Get("domElement"))

}
