package process

import (
	"fmt"

	"github.com/jasontconnell/go-project/data"
)

func WriteHello(msg data.Message) {
	fmt.Println(msg.Text)
}
