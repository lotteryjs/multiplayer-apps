package main

import (
	"fmt"

	"github.com/lonng/nano"
)

type room struct {
	group *nano.Group
}

func init() {
	fmt.Println("logo~")
}

func main() {
	nano.Listen(
		":3575",
		nano.WithDebugMode(),
		nano.WithIsWebsocket(true),
		nano.WithWSPath("/game"),
	)
}
