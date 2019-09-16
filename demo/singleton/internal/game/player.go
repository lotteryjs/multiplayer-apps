package game

import (
	"github.com/lonng/nano/session"
)

type Player struct {
	uid  int64  // 用户ID
	head string // 头像地址
	name string // 玩家名字
	ip   string // ip地址
	sex  int    // 性别

	// 玩家数据
	session *session.Session
}

func newPlayer(s *session.Session, uid int64, name, head, ip string, sex int) *Player {
	p := &Player{
		uid:  uid,
		name: name,
		head: head,
		ip:   ip,
		sex:  sex,
	}

	p.bindSession(s)

	return p
}

func (p *Player) bindSession(s *session.Session) {
	p.session = s
	p.session.Set(kCurPlayer, p)
}
