package model

type Uuid struct {
	Id        int64
	UidInUse  int64  `xorm:"not null index BIGINT(20) default 0"`
	UidOrigin int64  `xorm:"not null BIGINT(20) default 0"`
	Appid     string `xorm:"not null index VARCHAR(32) default"`
	Uuid      string `xorm:"not null VARCHAR(64) default"`
}
