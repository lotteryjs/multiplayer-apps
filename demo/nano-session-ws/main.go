package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/serialize/json"
	"github.com/lonng/nano/session"
)

type (
	// define component
	Room struct {
		component.Base
		group *nano.Group
	}

	// protocol messages
	UserMessage struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	NewUser struct {
		Content string `json:"content"`
	}

	AllMembers struct {
		Members []int64 `json:"members"`
	}

	JoinResponse struct {
		Code   int    `json:"code"`
		Result string `json:"result"`
	}
)

func NewRoom() *Room {
	return &Room{
		group: nano.NewGroup("room"),
	}
}

func (r *Room) AfterInit() {
	session.Lifetime.OnClosed(func(s *session.Session) {
		r.group.Leave(s)
	})
}

// Join room
func (r *Room) Join(s *session.Session, msg []byte) error {
	s.Bind(s.ID()) // binding session uid
	s.Push("onMembers", &AllMembers{Members: r.group.Members()})
	// // notify others
	r.group.Broadcast("onNewUser", &NewUser{Content: fmt.Sprintf("New user: %d", s.ID())})
	// // new user join group
	// r.group.Add(s) // add session to group
	return s.Response(&JoinResponse{Result: "sucess"})
}

// Send message
func (r *Room) Message(s *session.Session, msg *UserMessage) error {
	return r.group.Broadcast("onMessage", msg)
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
		nano.WithIsWebsocket(true),
		nano.WithSerializer(json.NewSerializer()))
}
