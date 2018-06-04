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
	count      = 0
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

func render() {
	scene.Call("render")
	count++

	if count < 2 {
		println(cube)
		println("position", cube.Get("position"))
	}
}

func onModelLoaded(obj *js.Object) {
	println("onModelLoaded", obj)
}

func main() {
	var document = js.Global.Get("document")
	var babylon = js.Global.Get("BABYLON")

	var canvas = document.Call("getElementById", "renderCanvas")
	var engine = babylon.Get("Engine").New(canvas, true)

	scene = babylon.Get("Scene").New(engine)
	vec3 := babylon.Get("Vector3").New(0, 3, 0)
	camera = babylon.Get("ArcRotateCamera").New("Camera", 0, 0, 5, vec3, scene)
	camera.Call("setPosition", babylon.Get("Vector3").New(0, 0, 150))
	camera.Call("attachControl", canvas)

	light1Vector := babylon.Get("Vector3").New(0, 1, 0)
	babylon.Get("HemisphericLight").New("light1", light1Vector, scene)

	loader := babylon.Get("AssetsManager").New(scene)

	cube = loader.Call("addMeshTask", "block", "", "models/", "block.obj")
	cube.Set("onSuccess", onModelLoaded)

	loader.Call("load")

	engine.Call("runRenderLoop", render)
}
