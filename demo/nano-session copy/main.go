package main

import (
	"fmt"
	"log"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

type (
	Room struct {
		component.Base
	}
)

// lifecyle method
func (r *Room) Init() {
	log.Print("Room.Init")
}

// lifecyle method
func (r *Room) AfterInit() {
	log.Print("Room.AfterInit")
}

// lifecyle method
func (r *Room) BeforeShutdown() {
	log.Print("BeforeShutdown")
}

// lifecyle method
func (r *Room) Shutdown() {
	log.Print("Shutdown")
}

func (r *Room) Join(s *session.Session, raw []byte) error {
	fmt.Printf("%v", s)
	return nil
}

func main() {
	components := &component.Components{}
	components.Register(&Room{})

	nano.Listen(":2828", nano.WithComponents(components))
}
