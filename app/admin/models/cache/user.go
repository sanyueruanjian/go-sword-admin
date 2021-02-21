package cache

import (
	"encoding/json"
	"fmt"
	"strconv"

	"project/app/admin/models/bo"
	"project/common/global"
)

const (
	UserInfoKeyFore  = "user::userInfo:id:"
	UserRecordsFore  = "user::records:auth:"
	KeyUserJob       = "job::user:"
	KeyUserRole      = "role::user:"
	KeyUserMenu      = "menu::user:"
	KeyUserDept      = "dept::user:"
	KeyUserDataScope = "data::user:"
)

func SetUserCenterInfoCache(userInfo *bo.UserCenterInfoBo) error {
	userByte, err := json.Marshal(userInfo)
	if err != nil {
		return err
	}
	return global.Rdb.Set(UserInfoKeyFore+strconv.Itoa(userInfo.User.Id), userByte, 0).Err()
}

func SetUserRecordsCache(userRecords []*bo.RecordUser, userId int) error {
	userByte, err := json.Marshal(userRecords)
	if err != nil {
		return err
	}
	return global.Rdb.Set(UserRecordsFore+strconv.Itoa(userId), userByte, 0).Err()
}

func GetUserRecordsCache(userId int) (userRecords []*bo.RecordUser, err error) {
	idStr := strconv.Itoa(userId)
	var userListByte []byte
	userListByte, err = global.Rdb.Get(UserRecordsFore + idStr).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(userListByte, &userRecords)
	return
}

func DelAllUserRecordsCache() error {
	recordsUser := global.Rdb.Keys(UserRecordsFore + "*").Val()
	pipLine := global.Rdb.Pipeline()
	for _, key := range recordsUser {
		pipLine.Del(key)
	}
	_, err := pipLine.Exec()
	return err
}

func GetUserCenterCache(id int) (userInfo *bo.UserCenterInfoBo, err error) {
	idStr := strconv.Itoa(id)
	var userByte []byte
	userByte, err = global.Rdb.Get(UserInfoKeyFore + idStr).Bytes()
	if err != nil {
		return nil, err
	}
	tmp := new(bo.UserCenterInfoBo)
	err = json.Unmarshal(userByte, tmp)
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

func DelAllUserCenterCache() error {
	keys := global.Rdb.Keys(UserInfoKeyFore + "*").Val()
	pipLine := global.Rdb.Pipeline()
	for _, key := range keys {
		pipLine.Del(key)
	}
	_, err := pipLine.Exec()
	return err
}

func DelUserAboutCache(userId int) error {
	var userKeys []string
	userKeys = append(userKeys, KeyUserDataScope, KeyUserDept,
		KeyUserDept, KeyUserJob, KeyUserMenu, KeyUserRole)
	pipLine := global.Rdb.Pipeline()
	for _, userKey := range userKeys {
		pipLine.Del(fmt.Sprintf("%s%d", userKey, userId))
	}
	_, err := pipLine.Exec()
	return err
}

func DelUsersAboutCache(usersId []int) error {
	var userKeys []string
	userKeys = append(userKeys, KeyUserDataScope, KeyUserDept,
		KeyUserDept, KeyUserJob, KeyUserMenu, KeyUserRole)
	pipLine := global.Rdb.Pipeline()
	for _, userId := range usersId {
		for _, userKey := range userKeys {
			pipLine.Del(fmt.Sprintf("%s%d", userKey, userId))
		}
		_, err := pipLine.Exec()
		if err != nil {
			return err
		}
	}
	return nil
}

func DelAllUserMenuCache() error {
	pipLine := global.Rdb.Pipeline()
	var userKeys []string
	userKeys = global.Rdb.Keys(KeyUserMenu + "*").Val()
	for _, userKey := range userKeys {
		pipLine.Del(userKey)
	}
	_, err := pipLine.Exec()
	return err
}
