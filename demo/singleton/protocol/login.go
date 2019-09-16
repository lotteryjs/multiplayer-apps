package protocol

type LoginToGameServerResponse struct {
	Uid      int64  `json:"acId"`
	Nickname string `json:"nickname"`
	HeadUrl  string `json:"headURL"`
	Sex      int    `json:"sex"`
	FangKa   int    `json:"fangka"`
}

type LoginToGameServerRequest struct {
	Name    string `json:"name"`
	Uid     int64  `json:"uid"`
	HeadUrl string `json:"headUrl"`
	Sex     int    `json:"sex"` //[0]未知 [1]男 [2]女
	FangKa  int    `json:"fangka"`
	IP      string `json:"ip"`
}
