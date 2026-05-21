GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/parse cmd/parse/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/matches cmd/matches/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/server-wasm cmd/server-wasm/main.go

server:
	@make wasmjs
	go run -mod $(GOMOD) cmd/server-wasm/main.go

wasmjs:
	GOOS=js GOARCH=wasm \
		go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags wasmjs \
		-o wasm/parse-js.wasm \
		cmd/parse-wasm/main.go

wasip1:
	GOARCH=wasm GOOS=wasip1 \
		go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags wasip1 \
		-o wasm/parse-p1.wasm \
		./cmd/parse-wasi/main.go

wasip2:
	tinygo build -target wasip2 -o wasm/parse-p2.wasm ./cmd/parse-wasi/main.go
