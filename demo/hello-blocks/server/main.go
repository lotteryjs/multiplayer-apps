package main

import (
	"fmt"
	"strings"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/serialize/json"
	"github.com/lonng/nano/session"
)

type Room struct {
	Group *nano.Group
}

type RoomManager struct {
	component.Base
	Rooms map[int]*Room
}

type JoinResponse struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

func (r *RoomManager) Init() {
	fmt.Println("RoomManager~Init")
}

func (r *RoomManager) Join(s *session.Session, msg []byte) error {
	return s.Response(&JoinResponse{Code: 1, Result: "Join~~Sucessed~~"})
}

func init() {
	fmt.Println("logo~" + nano.VERSION)
}

func main() {

	components := &component.Components{}
	components.Register(
		&RoomManager{},
		component.WithName("room"),
		component.WithNameFunc(strings.ToLower),
	)

	nano.Listen(
		":3575",
		nano.WithDebugMode(),
		nano.WithIsWebsocket(true),
		nano.WithWSPath("/game"),
		nano.WithComponents(components),
		nano.WithSerializer(json.NewSerializer()),
	)
}
