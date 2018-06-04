package main

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	scene      *js.Object
	camera     *js.Object
	cubeOffset = 0.001
	count      = 0
	cubeMesh   *js.Object
)

const (
	fps = 1000 / 30
)

func okClicked() {
	val := js.Global.Get("commandInput").Get("value").Float()
	cubeOffset = val
	println("cubeOffset updated to", cubeOffset)
}

func render() {
	if cubeMesh != nil {
		rotation := cubeMesh.Get("rotation")

		cubeX := rotation.Get("x").Float()
		cubeY := rotation.Get("y").Float()

		cubeX += cubeOffset
		cubeY += cubeOffset

		rotation.Set("x", cubeX)
		rotation.Set("y", cubeY)
	}

	scene.Call("render")
}

func onModelLoaded(obj *js.Object) {
	loadedMeshes := obj.Get("loadedMeshes")
	cubeMesh = loadedMeshes.Index(0)
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

	cube := loader.Call("addMeshTask", "block", "", "models/", "block.obj")
	cube.Set("onSuccess", onModelLoaded)

	loader.Call("load")

	js.Global.Get("okButton").Call("addEventListener", "click", okClicked)

	engine.Call("runRenderLoop", render)
}
