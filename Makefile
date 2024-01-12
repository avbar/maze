build-wasm:
	GOOS=js GOARCH=wasm go build -o web/maze.wasm cmd/app/main.go
	cp ${GOROOT}/misc/wasm/wasm_exec.js web

run:
	go run cmd/app/main.go

run-wsl:
	GOOS=windows go run cmd/app/main.go