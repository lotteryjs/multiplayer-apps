package web

import (
	"fmt"

	"github.com/lotteryjs/golang-server-dev/demo/singleton/db"
	"github.com/lotteryjs/golang-server-dev/demo/singleton/pkg/whitelist"
	"github.com/spf13/viper"
)

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

func Startup() {
	// setup database
	closer := dbStartup()
	defer closer()

	// enable white list
	enableWhiteList()

	fmt.Println("Web Server Startup")
}
