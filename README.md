# go-simple-eval

Simple query language for running predicates. 
It has a simple syntax and is easy to use. It also contains a small set of built-in functions.

## Installation

```shell
go get -u github.com/stirante/go-simple-eval
```

## Usage

```go
eval.Eval("test > 2", map[string]interface{}{
	"test": 10,
})
```

## Development

### Prerequisites

- [Go v1.18 or later](https://golang.org/)

### Setup

```powershell
# Clone the repository
git clone https://github.com/stirante/go-simple-eval
cd go-simple-eval
# Install dependencies
go mod vendor
cd scripts
# Setup ANTLR
./setup_env.ps1
# Compile grammar
./compile_antlr.ps1
# Go back to the root directory
cd ..
```

### Building

```powershell
go build github.com/stirante/go-simple-eval
```

### Notes

All changes to `grammar/SimpleQueryLanguage.g4` must be compiled with `scripts/compile_antlr.ps1` before they can be used.
Due to the issue with generating the parser in Go, each new rule should be added to method `visit` in `query/expression_visitor.go` manually.
