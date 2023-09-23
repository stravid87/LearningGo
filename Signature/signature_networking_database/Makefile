client_first: cmd/client_first.go
	go run $<

client_second: cmd/client_second.go
	go run $<

server: cmd/server.go
	go run $<

copy_assets:
	$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./assets

update_wasm:
	GOOS=js GOARCH=wasm go build --o assets/main.wasm cmd/wasm/*

