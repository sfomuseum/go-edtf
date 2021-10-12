cli:
	go build -mod vendor -o bin/parse cmd/parse/main.go
	go build -mod vendor -o bin/matches cmd/matches/main.go
