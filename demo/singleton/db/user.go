package db

import (
	"time"

	"github.com/lotteryjs/golang-server-dev/demo/singleton/db/model"
	"github.com/lotteryjs/golang-server-dev/demo/singleton/pkg/errutil"
	"github.com/lotteryjs/golang-server-dev/demo/singleton/protocol"
)

func QueryGuestUser(appId string, imei string) (*model.User, error) {
	bean := &model.Register{
		Imei:  imei,
		AppId: appId,
	}

	ok, err := database.Get(bean)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errutil.ErrUserNotFound
	}

	user := &model.User{
		Id: bean.Uid,
	}
	ok, err = database.Get(user)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errutil.ErrUserNotFound
	}

	return user, nil
}

//InsertUser insert a new user
func InsertUser(u *model.User) error {
	if u == nil {
		return nil
	}
	_, err := database.Insert(u)
	return err
}

func InsertLoginLog(uid int64, d protocol.Device, appid string, channelID string) {
	// Insert user operation record
	log := &model.Login{
		Uid:       uid,
		Remote:    d.Remote,
		Ip:        d.IP,
		Imei:      d.IMEI,
		Os:        d.OS,
		Model:     d.Model,
		AppId:     appid,
		ChannelId: channelID,
		LoginAt:   time.Now().Unix(),
	}
	userOnline(uid)
	chWrite <- log
}

func RegisterUserLog(u *model.User, d protocol.Device, appid string, channelID string, regType int) {
	// Insert user register record
	t := time.Now().Unix()
	reg := &model.Register{
		Uid:          u.Id,
		Remote:       d.Remote,
		Ip:           d.IP,
		Imei:         d.IMEI,
		Os:           d.OS,
		Model:        d.Model,
		AppId:        appid,
		ChannelId:    channelID,
		RegisterAt:   t,
		RegisterType: regType,
	}
	InsertRegister(reg)
}

func InsertRegister(reg *model.Register) {
	chWrite <- reg
}

func userOnline(uid int64) error {
	u := &model.User{IsOnline: UserOnline, LastLoginAt: time.Now().Unix()}
	if _, err := database.Where("id=?", uid).Update(u); err != nil {
		return err
	}
	return nil
}
