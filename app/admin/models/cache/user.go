package cache

import (
	"encoding/json"
	"project/app/admin/models/bo"
	"project/common/global"
	"project/utils/config"
	"strconv"
	"time"
)

const (
	UserInfoKeyFore = "user::userInfo:id"
)

func SetUserCenterListCache(userInfo *bo.UserCenterInfoBo) error {
	userByte, err := json.Marshal(userInfo)
	if err != nil {
		return err
	}
	return global.Rdb.Set(UserInfoKeyFore+strconv.Itoa(userInfo.User.Id), userByte, time.Duration(config.JwtConfig.Timeout)*time.Second).Err()
}

func GetUserCenterCache(id int) (userInfo *bo.UserCenterInfoBo, err error) {
	idStr := strconv.Itoa(id)
	var userByte []byte
	userByte, err = global.Rdb.Get(UserInfoKeyFore + idStr).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(userByte, userInfo)
	return userInfo, err
}

func DelUserCenterCache(id int) error {
	idStr := strconv.Itoa(id)
	return global.Rdb.Del(UserInfoKeyFore + idStr).Err()
}

func DelManyUserCenterCache(ids []int) error {
	pipLine := global.Rdb.Pipeline()
	for _, id := range ids {
		idStr := strconv.Itoa(id)
		key := UserInfoKeyFore + idStr
		pipLine.Del(key)
	}
	_, err := pipLine.Exec()
	return err
}
