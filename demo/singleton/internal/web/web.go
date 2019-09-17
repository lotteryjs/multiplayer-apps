package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lonng/nex"
	"github.com/lotteryjs/golang-server-dev/demo/singleton/db"
	"github.com/lotteryjs/golang-server-dev/demo/singleton/internal/web/api"
	"github.com/lotteryjs/golang-server-dev/demo/singleton/pkg/algoutil"
	"github.com/lotteryjs/golang-server-dev/demo/singleton/pkg/whitelist"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var logger = log.WithField("component", "http")

func dbStartup() func() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.dbname"),
		viper.GetString("database.args"))

	return db.MustStartup(
		dsn,
		db.MaxIdleConns(viper.GetInt("database.max_idle_conns")),
		db.MaxIdleConns(viper.GetInt("database.max_open_conns")),
		db.ShowSQL(viper.GetBool("database.show_sql")))
}

func enableWhiteList() {
	whitelist.Setup(viper.GetStringSlice("whitelist.ip"))
}

func pongHandler() (string, error) {
	return "pong", nil
}

func logRequest(ctx context.Context, r *http.Request) (context.Context, error) {
	if uri := r.RequestURI; uri != "/ping" {
		logger.Debugf("Method=%s, RemoteAddr=%s URL=%s", r.Method, r.RemoteAddr, uri)
	}
	return ctx, nil
}

func startupService() http.Handler {
	var (
		mux    = http.NewServeMux()
		webDir = viper.GetString("webserver.static_dir")
	)

	nex.Before(logRequest)
	mux.Handle("/v1/user/", api.MakeLoginService())

	mux.Handle("/v1/test", nex.Handler(
		func() (string, error) {
			return "test", nil
		}),
	)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(webDir))))
	mux.Handle("/ping", nex.Handler(pongHandler))

	return algoutil.AccessControl(algoutil.OptionControl(mux))
}

// Startup for Web Server
func Startup() {
	// setup database
	closer := dbStartup()
	defer closer()

	// enable white list
	enableWhiteList()

	var (
		addr      = viper.GetString("webserver.addr")
		cert      = viper.GetString("webserver.certificates.cert")
		key       = viper.GetString("webserver.certificates.key")
		enableSSL = viper.GetBool("webserver.enable_ssl")
	)

	logger.Infof("Web service addr: %s(enable ssl: %v)", addr, enableSSL)
	go func() {
		// http service
		mux := startupService()
		if enableSSL {
			log.Fatal(http.ListenAndServeTLS(addr, cert, key, mux))
		} else {
			log.Fatal(http.ListenAndServe(addr, mux))
		}
	}()

	// stop server
	sg := make(chan os.Signal)
	signal.Notify(sg, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	select {
	case s := <-sg:
		log.Infof("got signal: %s", s.String())
	}
}
