# webgl

[![GoDoc](https://godoc.org/github.com/justinclift/webgl?status.svg)](https://godoc.org/github.com/justinclift/webgl)
[![Go Report Card](https://goreportcard.com/badge/github.com/justinclift/webgl)](https://goreportcard.com/report/github.com/justinclift/webgl)

[TinyGo](https://github.com/tinygo-org/tinygo) bindings for [WebGL 1.0](https://www.khronos.org/registry/webgl/specs/latest/1.0/) context.

## Example

![Screenshot](https://cloud.githubusercontent.com/assets/1924134/3566022/5d81f2d0-0ae0-11e4-82e4-3cb33b83d8d3.png)

webgl_example.go:

```Go
package main

import (
	"syscall/js"

	"github.com/justinclift/webgl"
)

func main() {
	// Get the canvas element
	doc := js.Global().Get("document")
	canvas := doc.Call("getElementById", "mycanvas")
	width := canvas.Get("clientWidth").Int()
	height := canvas.Get("clientHeight").Int()
	canvas.Call("setAttribute", "width", width)
	canvas.Call("setAttribute", "height", height)

	// Set the desired WebGL context attributes
	attrs := webgl.DefaultAttributes()
	attrs.Alpha = false

	// Create the WebGL context
	gl, err := webgl.NewContext(&canvas, attrs)
	if err != nil {
		js.Global().Call("alert", "Error: "+err.Error())
		return
	}

	// Do something with the WebGL context
	gl.ClearColor(0.8, 0.3, 0.01, 1)
	gl.Clear(webgl.COLOR_BUFFER_BIT)
}
```

index.html:

```html
<!DOCTYPE html>

<html>
<head>
	<meta charset="utf-8" />
	<title>TinyGo wasm WebGL Fundamentals</title>
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<script src="wasm_exec.js" defer></script>
	<script src="wasm.js" defer></script>
	<style>
		body,pre { margin:0;padding:0; }
		#mycanvas {
			position:fixed;
			opacity:0.9;
			width: 100%;
			height:100%;
			top:0;right:0;bottom:0;left:0;
		}
	</style>
</head>
<body>
	<canvas id="mycanvas">Your browser doesn't appear to support the canvas tag.</canvas>
</body>
</html>

```

wasm.js:

```js
'use strict';

const WASM_URL = 'wasm.wasm';
var wasm;

// Load and run the wasm
function init() {
  const go = new Go();
  if ('instantiateStreaming' in WebAssembly) {
    WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(function (obj) {
      wasm = obj.instance;
      go.run(wasm); // Initial setup
    })
  } else {
    fetch(WASM_URL).then(resp =>
      resp.arrayBuffer()
    ).then(bytes =>
      WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
        wasm = obj.instance;
        go.run(wasm);
      })
    )
  }
}

init();

```

wasm_exec.js should be obtained from the TinyGo source repository.
