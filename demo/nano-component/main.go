package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

type Demo struct{}

func (r *Demo) Init() {
	log.Print("Demo.Init")
}

func (r *Demo) AfterInit() {
	log.Print("Demo.AfterInit")
}

func (r *Demo) BeforeShutdown() {
	log.Print("BeforeShutdown")
}

func (r *Demo) Shutdown() {
	log.Print("Shutdown")
}

// 以下的Handler不会自动将消息反序列化，会将客户端发送过来的消息直接当作参数传进来
func (r *Demo) Join(s *session.Session, raw []byte) error {
	fmt.Printf("%v", s)
	return nil
}

func main() {
	components := &component.Components{}
	// components.Register(&Demo{})
	components.Register(&Demo{}, component.WithName("demo"), component.WithNameFunc(strings.ToLower))

	nano.Listen(":2828", nano.WithComponents(components))
}
