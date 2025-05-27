GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/parse cmd/parse/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/matches cmd/matches/main.go

wasmjs:
	GOOS=js GOARCH=wasm \
		go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags wasmjs \
		-o wasm/parse-js.wasm \
		cmd/parse-wasm/main.go

wasip:
	GOARCH=wasm GOOS=wasip1 \
		go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags wasip1 \
		-o wasm/parse-p1.wasm \
		./cmd/parse-wasi/main.go

