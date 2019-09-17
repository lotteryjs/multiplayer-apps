package game

import (
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
	"github.com/lotteryjs/golang-server-dev/demo/singleton/protocol"
	log "github.com/sirupsen/logrus"
)

const kickResetBacklog = 8

var defaultManager = NewManager()

type (
	Manager struct {
		component.Base
		group   *nano.Group       // 广播channel
		players map[int64]*Player // 所有的玩家
		chKick  chan int64        // 退出队列
		chReset chan int64        // 重置队列
	}
)

func NewManager() *Manager {
	return &Manager{
		group:   nano.NewGroup("_SYSTEM_MESSAGE_BROADCAST"),
		players: map[int64]*Player{},
		chKick:  make(chan int64, kickResetBacklog),
		chReset: make(chan int64, kickResetBacklog),
	}
}

func (m *Manager) AfterInit() {
	session.Lifetime.OnClosed(func(s *session.Session) {
		m.group.Leave(s)
	})
}

func (m *Manager) Login(s *session.Session, req *protocol.LoginToGameServerRequest) error {
	uid := req.Uid
	s.Bind(uid)

	log.Infof("玩家: %d登录: %+v", uid, req)
	if p, ok := m.player(uid); !ok {
		log.Infof("玩家: %d不在线，创建新的玩家", uid)
		p = newPlayer(s, uid, req.Name, req.HeadUrl, req.IP, req.Sex)
		m.setPlayer(uid, p)
	} else {
		log.Infof("玩家: %d已经在线", uid)
		//ToDo...
	}

	// 添加到广播频道
	m.group.Add(s)

	res := &protocol.LoginToGameServerResponse{
		Uid:      s.UID(),
		Nickname: req.Name,
		Sex:      req.Sex,
		HeadUrl:  req.HeadUrl,
	}

	return s.Response(res)
}

func (m *Manager) player(uid int64) (*Player, bool) {
	p, ok := m.players[uid]

	return p, ok
}

func (m *Manager) setPlayer(uid int64, p *Player) {
	if _, ok := m.players[uid]; ok {
		log.Warnf("玩家已经存在，正在覆盖玩家， UID=%d", uid)
	}
	m.players[uid] = p
}

func (m *Manager) sessionCount() int {
	return len(m.players)
}

func (m *Manager) offline(uid int64) {
	delete(m.players, uid)
	log.Infof("玩家: %d从在线列表中删除, 剩余：%d", uid, len(m.players))
}
