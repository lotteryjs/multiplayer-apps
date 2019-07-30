package main

import (
	"fmt"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

type (
	Room struct {
		component.Base
	}
)

func (r *Room) Join(s *session.Session, raw []byte) error {
	fmt.Printf("%v", s)
	return nil
}

func main() {
	components := &component.Components{}
	components.Register(&Room{})

	nano.Listen(":2828", nano.WithComponents(components))
}
