//go:build wasmjs
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"syscall/js"

	"github.com/sfomuseum/go-edtf/parser"	
)

func ParseFunc() js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		edtf_str := args[0].String()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			edtf_d, err := parser.ParseString(edtf_str)

			if err != nil {
				reject.Invoke(fmt.Printf("Failed to parse '%s', %v\n", edtf_str, err))
				return nil
			}

			enc, err := json.Marshal(edtf_d)

			if err != nil {
				reject.Invoke(fmt.Printf("Failed to marshal result for '%s', %v\n", edtf_str, err))
				return nil
			}

			resolve.Invoke(string(enc))
			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func main() {

	parse_func := ParseFunc()
	defer parse_func.Release()

	js.Global().Set("parse_edtf", parse_func)

	c := make(chan struct{}, 0)

	log.Println("WASM EDTF parser initialized")
	<-c
}
