package cache

import (
	"project/common/global"
	"project/utils/config"
	"strconv"
	"time"
)

func GetAllMenuIdCacheKeys() []string {
	return global.Rdb.Keys("menu::id:*").Val()
}

//查询所有菜单缓存
func GetAllMenuCache(menuIdInRedis []string) (data [][]byte, err error) {
	result, err := global.Rdb.MGet(menuIdInRedis...).Result()
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		data = append(data, []byte(v.(string)))
	}
	return data, nil
}

//做所有菜单缓存
func SetMenuCache(allMenu map[string][]byte) error {
	pipLine := global.Rdb.Pipeline()
	for k, v := range allMenu {
		pipLine.Set(k, v, time.Duration(config.JwtConfig.Timeout)*time.Second)
	}
	if _, err := pipLine.Exec(); err != nil {
		return err
	}
	return nil
}

//查询所有菜单bo缓存
func SetMenuListCache(v []byte, k string) error {
	return global.Rdb.Set(k, v, time.Duration(config.JwtConfig.Timeout)*time.Second).Err()
}

//查询所有菜单bo缓存
func GetMenuListCache(k string) ([]byte, error) {
	return global.Rdb.Get(k).Bytes()
}

//根据id查询缓存
func GetMenuByIdCache(id int) (menu []byte, err error) {
	strId := strconv.Itoa(id)
	return global.Rdb.Get("menu::id:" + strId).Bytes()
}

//根据id删除缓存
func DeleteMenuByIdCache(ids []int) error {
	pipLine := global.Rdb.Pipeline()
	for _, id := range ids {
		strId := strconv.Itoa(id)
		pipLine.Del("menu::id:" + strId)
	}
	_, err := pipLine.Exec()
	return err
}

func DeleteAllMenuIdCache() error {
	keys := GetAllMenuIdCacheKeys()
	pipLine := global.Rdb.Pipeline()
	for _, key := range keys {
		pipLine.Del(key)
	}
	_, err := pipLine.Exec()
	return err
}

func DeleteAllUserMenuCache(keys []string) error {
	pipLine := global.Rdb.Pipeline()
	for _, key := range keys {
		pipLine.Del(key)
	}
	_, err := pipLine.Exec()
	return err
}
