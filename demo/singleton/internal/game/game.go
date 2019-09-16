package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/serialize/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	logger = log.WithField("component", "game")
)

func Startup() {
	rand.Seed(time.Now().Unix())

	heartbeat := viper.GetInt("core.heartbeat")
	if heartbeat < 5 {
		heartbeat = 5
	}

	logger.Infof("当前心跳时间间隔: %d秒", heartbeat)
	logger.Info("game service starup")

	// register game handler
	comps := &component.Components{}
	comps.Register(defaultManager)

	// 加密管道
	// c := newCrypto()
	// pip := pipeline.New()
	// pip.Inbound().PushBack(c.inbound)
	// pip.Outbound().PushBack(c.outbound)

	addr := fmt.Sprintf(":%d", viper.GetInt("game-server.port"))
	nano.Listen(addr,
		// nano.WithPipeline(pip),
		nano.WithHeartbeatInterval(time.Duration(heartbeat)*time.Second),
		nano.WithLogger(log.WithField("component", "nano")),
		nano.WithDebugMode(),
		// nano.WithIsWebsocket(true),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithComponents(comps),
	)
}
