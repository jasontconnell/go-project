package test

import (
	"testing"

	"github.com/jasontconnell/go-project/data"
	"github.com/jasontconnell/go-project/process"
)

func TestLib(t *testing.T) {
	process.WriteHello(data.Message{Text: "Hello World!"})

}
