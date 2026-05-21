package wasm

import (
	"embed"
)

//go:embed *.wasm
var FS embed.FS
