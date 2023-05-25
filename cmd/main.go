package main

import (
	"github.com/jasontconnell/go-project/data"
	"github.com/jasontconnell/go-project/process"
)

func main() {
	process.WriteHello(data.Message{Text: "Hello World!"})
}
