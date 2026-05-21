package www

import (
	"embed"
)

//go:embed *.html css/* javascript/*
var FS embed.FS
