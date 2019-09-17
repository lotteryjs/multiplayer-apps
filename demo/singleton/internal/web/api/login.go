package api

import (
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/gorilla/mux"
	"github.com/lonng/nex"
	"github.com/lotteryjs/golang-server-dev/demo/singleton/db"
	"github.com/lotteryjs/golang-server-dev/demo/singleton/db/model"
	"github.com/lotteryjs/golang-server-dev/demo/singleton/protocol"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	host     string                // 服务器地址
	port     int                   // 服务器端口
	config   protocol.ClientConfig // 远程配置
	messages []string              // 广播消息
	logger   = log.WithFields(log.Fields{"component": "http", "service": "login"})

	// 游客登陆
	enableGuest   = false
	guestChannels = []string{}

	enableDebug = false
)

func AddMessage(message string) {
	messages = append(messages, message)
}

func MakeLoginService() http.Handler {
	host = viper.GetString("game-server.host")
	port = viper.GetInt("game-server.port")

	// 更新相关配置
	config.Heartbeat = viper.GetInt("core.heartbeat")

	// 游客相关配置
	enableGuest = viper.GetBool("login.guest")
	guestChannels = viper.GetStringSlice("login.lists")
	logger.Infof("是否开启游客登陆: %t, 渠道列表: %v", enableGuest, guestChannels)

	if config.Heartbeat < 5 {
		config.Heartbeat = 5
	}

	messages = viper.GetStringSlice("broadcast.message")

	logger.Debugf("version infomation: %+v", config)
	logger.Debugf("广播消息: %v", messages)

	router := mux.NewRouter()
	router.Handle("/v1/user/login/query", nex.Handler(queryHandler)).Methods("POST")      //三方登录
	router.Handle("/v1/user/login/guest", nex.Handler(guestLoginHandler)).Methods("POST") //三方登录
	return router
}

func ip(addr string) string {
	addr = strings.TrimSpace(addr)
	deflt := "127.0.0.1"
	if addr == "" {
		return deflt
	}

	if parts := strings.Split(addr, ":"); len(parts) > 0 {
		return parts[0]
	}

	return deflt
}

// 过滤emoji表情, 设置为*
func filterEmoji(content string) string {
	new_content := ""
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			new_content += string(value)
		} else {
			new_content += "*"
		}
	}
	return new_content
}

func checkSession(uid int64) {
	// fixed: 之前已有session未断开
	// 检查是否该玩家是否有未断开的网络连接, 把之前的号顶掉
	// game.Kick(uid)
}

func guestLoginHandler(r *http.Request, data *protocol.LoginRequest) (*protocol.LoginResponse, error) {
	data.Device.IMEI = data.IMEI

	logger.Infof("游客登录IEMEI: %s", data.Device.IMEI)

	user, err := db.QueryGuestUser(data.AppID, data.Device.IMEI)
	if err != nil {
		// 生成一个新用户
		user = &model.User{
			Status:   db.StatusNormal,
			IsOnline: db.UserOffline,
			Role:     db.RoleTypeThird,
		}

		if err := db.InsertUser(user); err != nil {
			logger.Error(err.Error())
			return nil, err
		}

		db.RegisterUserLog(user, data.Device, data.AppID, data.ChannelID, protocol.RegTypeThird) //注册记录
	}

	checkSession(user.Id)

	resp := &protocol.LoginResponse{
		Uid:      user.Id,
		HeadUrl:  "http://wx.qlogo.cn/mmopen/s962LEwpLxhQSOnarDnceXjSxVGaibMRsvRM4EIWic0U6fQdkpqz4Vr8XS8D81QKfyYuwjwm2M2ibsFY8mia8ic51ww/0",
		Sex:      1,
		IP:       host,
		Port:     port,
		PlayerIP: ip(r.RemoteAddr),
		Config:   config,
		Messages: messages,
		Debug:    0, //user.Debug,
	}
	resp.Name = fmt.Sprintf("G%d", resp.Uid)

	// 插入登陆记录
	device := protocol.Device{
		IP:     ip(r.RemoteAddr),
		Remote: r.RemoteAddr,
	}
	db.InsertLoginLog(user.Id, device, data.AppID, data.ChannelID)

	return resp, nil
}

// 查询是否使用游客登陆
type (
	queryRequest struct {
		AppId     string `json:"appId"`
		ChannelId string `json:"channelId"`
	}

	queryResponse struct {
		Code  int  `json:"code"`
		Guest bool `json:"guest"`
	}
)

var (
	forbidGuest  = &queryResponse{Guest: false}
	accepetGuest = &queryResponse{Guest: true}
)

func queryHandler(query *queryRequest) (*queryResponse, error) {
	logger.Infof("%v", query)
	if !enableGuest {
		return forbidGuest, nil
	}

	for _, s := range guestChannels {
		if query.ChannelId == s {
			return accepetGuest, nil
		}
	}

	return forbidGuest, nil
}
