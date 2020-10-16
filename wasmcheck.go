// +build !wasm

package webgl

func init() {
	panic("This package must be built targeting WASM")
}
