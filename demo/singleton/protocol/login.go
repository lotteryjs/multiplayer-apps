package protocol

type ClientConfig struct {
	Heartbeat int `json:"heartbeat"`

	AppId  string `json:"appId"`
	AppKey string `json:"appKey"`
}

type LoginToGameServerResponse struct {
	Uid      int64  `json:"acId"`
	Nickname string `json:"nickname"`
	HeadUrl  string `json:"headURL"`
	Sex      int    `json:"sex"`
}

type LoginToGameServerRequest struct {
	Name    string `json:"name"`
	Uid     int64  `json:"uid"`
	HeadUrl string `json:"headUrl"`
	Sex     int    `json:"sex"` //[0]未知 [1]男 [2]女
	IP      string `json:"ip"`
}

type LoginRequest struct {
	AppID     string `json:"appId"`     //用户来自于哪一个应用
	ChannelID string `json:"channelId"` //用户来自于哪一个渠道
	IMEI      string `json:"imei"`
	Device    Device `json:"device"`
}

type LoginResponse struct {
	Code     int          `json:"code"`
	Name     string       `json:"name"`
	Uid      int64        `json:"uid"`
	HeadUrl  string       `json:"headUrl"`
	Sex      int          `json:"sex"` //[0]未知 [1]男 [2]女
	IP       string       `json:"ip"`
	Port     int          `json:"port"`
	PlayerIP string       `json:"playerIp"`
	Config   ClientConfig `json:"config"`
	Messages []string     `json:"messages"`
	Debug    int          `json:"debug"`
}
