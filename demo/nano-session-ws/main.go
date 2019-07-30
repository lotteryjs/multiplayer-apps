package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

type Room struct {
	component.Base
}

func (r *Room) Join(s *session.Session, raw []byte) error {
	fmt.Printf("%s", s)
	return nil
}

func main() {
	components := &component.Components{}

	components.Register(&Room{},
		component.WithName("room"),
		component.WithNameFunc(strings.ToLower))

	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("../web"))))
	nano.Listen(":2828",
		nano.WithComponents(components),
		nano.WithWSPath("/ws"),
		nano.WithIsWebsocket(true))
}
