package game

import (
	"github.com/lonng/nano/session"
	"github.com/lotteryjs/golang-server-dev/demo/singleton/pkg/errutil"
)

func playerWithSession(s *session.Session) (*Player, error) {
	p, ok := s.Value(kCurPlayer).(*Player)
	if !ok {
		return nil, errutil.ErrPlayerNotFound
	}
	return p, nil
}
