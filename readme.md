## WASM
You can compile it with similar commands:

[//]: # (https://github.com/golang/go/wiki/WebAssembly#getting-started)

```cli
GOOS=js GOARCH=wasm go build -o main.wasm

cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```
### Run with GOEXEC
install goexec: go get -u github.com/shurcooL/goexec
```go
goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
```
Run without adding to PATH
```go
GOOS=js GOARCH=wasm go run -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" .
```
### TinyGO
Build with TinyGO
```go
go build -o hello ./hello.go

tinygo build -o hello ./hello.go
```