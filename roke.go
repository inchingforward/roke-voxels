package main

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	cube       *js.Object
	scene      *js.Object
	camera     *js.Object
	renderer   *js.Object
	cubeOffset = 0.05
)

const (
	fps = 1000 / 30
)

func renderLoop() {
	rot := cube.Get("rotation")

	cubeX := rot.Get("x").Float()
	cubeY := rot.Get("y").Float()

	cubeX += cubeOffset
	cubeY += cubeOffset

	rot.Set("x", cubeX)
	rot.Set("y", cubeY)

	renderer.Call("render", scene, camera)

	js.Global.Call("setTimeout", renderLoop, fps)
}

func okClicked() {
	val := js.Global.Get("commandInput").Get("value").Float()
	cubeOffset = val
}

func main() {
	var document = js.Global.Get("document")
	var body = document.Get("body")

	var three = js.Global.Get("THREE")
	scene = three.Get("Scene").New()
	camera = three.Get("PerspectiveCamera").New(70, 1, 0.1, 1000)

	var boxGeo = three.Get("BoxGeometry").New(1, 1, 1)
	var material = three.Get("MeshBasicMaterial").New(map[string]int{"color": 0x0066cc})
	cube = three.Get("Mesh").New(boxGeo, material)

	scene.Call("add", cube)
	var pos = camera.Get("position")
	pos.Set("z", 5)

	renderer = three.Get("WebGLRenderer").New()

	renderer.Call("setSize", 500, 500)

	body.Call("appendChild", renderer.Get("domElement"))

	js.Global.Get("okButton").Call("addEventListener", "click", okClicked)

	renderLoop()
}
