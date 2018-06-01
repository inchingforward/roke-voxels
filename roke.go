package main

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	cube     *js.Object
	cubeX    float64
	cubeY    float64
	scene    *js.Object
	camera   *js.Object
	renderer *js.Object
)

func renderLoop() {
	cubeX += 0.05
	cubeY += 0.05

	rot := cube.Get("rotation")
	rot.Set("x", cubeX)
	rot.Set("y", cubeY)

	renderer.Call("render", scene, camera)

	js.Global.Call("setTimeout", renderLoop, 1000/30)
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

	renderLoop()
}
