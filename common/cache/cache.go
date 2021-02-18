package cache

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"

	"project/common/global"
)

const DelUserCacheError = "删除用户缓存失败"

func DelUserCacheById(key string, ids *[]int) (err error) {
	result := new([]*redis.IntCmd)
	pipe := global.Rdb.TxPipeline()
	for _, id := range *ids {
		*result = append(*result, pipe.Del(fmt.Sprintf("%s%d", key, id)))
	}
	_, _ = pipe.Exec()

	for _, r := range *result {
		if r.Err() != nil {
			return errors.New(DelUserCacheError)
		}
	}
	return
}
