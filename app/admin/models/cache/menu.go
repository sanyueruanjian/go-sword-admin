package cache

import (
	"project/common/global"
	"strconv"
)

const (
	MenuIdKeyFore = "menu::id:"
)

func GetAllMenuIdCacheKeys() []string {
	return global.Rdb.Keys(MenuIdKeyFore + "*").Val()
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
		pipLine.Set(k, v, 0)
	}
	if _, err := pipLine.Exec(); err != nil {
		return err
	}
	return nil
}

//设置菜单bo缓存
func SetMenuListCache(v []byte, pid int) error {
	return global.Rdb.Set("menu::pid:"+strconv.Itoa(pid), v, 0).Err()
}

//查询所有菜单bo缓存
func GetMenuListCache(k string) ([]byte, error) {
	return global.Rdb.Get(k).Bytes()
}

//删除菜单bo缓存
func DelMenuListCache(pIDs []int) error {
	pipeline := global.Rdb.Pipeline()
	for _, pid := range pIDs {
		pipeline.Del("menu::pid:" + strconv.Itoa(pid))
	}
	_, err := pipeline.Exec()
	return err
}

//根据id查询缓存
func GetMenuByIdCache(id int) (menu []byte, err error) {
	strId := strconv.Itoa(id)
	return global.Rdb.Get(MenuIdKeyFore + strId).Bytes()
}

//根据id删除缓存
func DeleteMenuByIdCache(ids []int) error {
	pipLine := global.Rdb.Pipeline()
	for _, id := range ids {
		strId := strconv.Itoa(id)
		pipLine.Del(MenuIdKeyFore + strId)
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

func DeleteAllUserNeedMenuCache() error {
	keys := global.Rdb.Keys("menu::userNeed:").Val()
	pipLine := global.Rdb.Pipeline()
	for _, key := range keys {
		pipLine.Del(key)
	}
	_, err := pipLine.Exec()
	return err
}
