package model

type User struct {
	Id          int64
	Algo        string `xorm:"not null VARCHAR(16) default"`
	Hash        string `xorm:"not null VARCHAR(64) default"`
	Salt        string `xorm:"not null VARCHAR(64) default"`
	Role        int    `xorm:"not null TINYINT(3) default 1"`
	Status      int    `xorm:"not null TINYINT(3) default 1"`
	IsOnline    int    `xorm:"not null TINYINT(1) default 1"`
	LastLoginAt int64  `xorm:"not null index BIGINT(11) default"`
	PrivKey     string `xorm:"not null VARCHAR(512) default"`
	PubKey      string `xorm:"not null VARCHAR(128) default"`
	RegisterAt  int64  `xorm:"not null index BIGINT(20) default 0"`
	Debug       int    `xorm:"not null index TINYINT(1) default 0"`
}

type Register struct {
	Id           int64
	Uid          int64  `xorm:"not null index BIGINT(20) default"`
	Remote       string `xorm:"not null VARCHAR(40) default"`
	Ip           string `xorm:"not null VARCHAR(40) default"`
	Imei         string `xorm:"not null VARCHAR(128) default"`
	Os           string `xorm:"not null VARCHAR(20) default"`
	Model        string `xorm:"not null VARCHAR(20) default"`
	AppId        string `xorm:"not null index VARCHAR(32) default"`
	ChannelId    string `xorm:"not null index VARCHAR(32) default"`
	RegisterAt   int64  `xorm:"not null index BIGINT(11) default"`
	RegisterType int    `xorm:"not null index TINYINT(8) default"`
}

type Login struct {
	Id        int64
	Uid       int64  `xorm:"not null index BIGINT(20) default"`
	Remote    string `xorm:"not null VARCHAR(40) default"`
	Ip        string `xorm:"not null VARCHAR(40) default"`
	Model     string `xorm:"not null VARCHAR(64) default"`
	Imei      string `xorm:"not null VARCHAR(32) default"`
	Os        string `xorm:"not null VARCHAR(64) default"`
	AppId     string `xorm:"not null VARCHAR(64) default"`
	ChannelId string `xorm:"not null VARCHAR(32) default"`
	LoginAt   int64  `xorm:"not null BIGINT(11) default"`
	LogoutAt  int64  `xorm:"not null BIGINT(11) default"`
}
