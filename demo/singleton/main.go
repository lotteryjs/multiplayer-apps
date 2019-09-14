package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/lotteryjs/golang-server-dev/demo/singleton/internal/game"
	"github.com/lotteryjs/golang-server-dev/demo/singleton/internal/web"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Singleton Mode Server"
	app.Author = "Singleton"
	app.Version = "0.0.1"
	app.Copyright = "singleton team reserved"
	app.Usage = "singleton server"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "./configs/config.toml",
			Usage: "load configuration from `FILE`",
		},
		cli.BoolFlag{
			Name:  "cpuprofile",
			Usage: "enable cpu profile",
		},
	}

	app.Action = serve
	app.Run(os.Args)
}

func serve(c *cli.Context) error {
	viper.SetConfigType("toml")
	viper.SetConfigFile(c.String("config"))
	viper.ReadInConfig()

	log.SetFormatter(&log.TextFormatter{DisableColors: true})
	if viper.GetBool("core.debug") {
		log.SetLevel(log.DebugLevel)
	}

	if c.Bool("cpuprofile") {
		filename := fmt.Sprintf("cpuprofile-%d.pprof", time.Now().Unix())
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() { defer wg.Done(); game.Startup() }() // game server
	go func() { defer wg.Done(); web.Startup() }()  // web server

	wg.Wait()
	return nil
}
