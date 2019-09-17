package db

const (
	defaultMaxConns = 10
)

// Users表中role字段的取值
const (
	RoleTypeAdmin = 1 //管理员账号
	RoleTypeThird = 2 //三方平台账号
)

const (
	StatusNormal  = 1 //正常
	StatusDeleted = 2 //删除
	StatusFreezed = 3 //冻结
	StatusBound   = 4 //绑定
)

const (
	UserOffline = 1 //离线
	UserOnline  = 2 //在线
)
