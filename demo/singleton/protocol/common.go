package protocol

const (
	RegTypeThird = 5 //三方平台添加账号
)

type Device struct {
	IMEI   string `json:"imei"`   //设备的imei号
	OS     string `json:"os"`     //os版本号
	Model  string `json:"model"`  //硬件型号
	IP     string `json:"ip"`     //内网IP
	Remote string `json:"remote"` //外网IP
}
