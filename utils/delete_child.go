package utils

import "project/common/global"

//删除子结构 辅助递归删除
func DeleteChild(tableName string, id int) (err error) {
	err = global.Eloquent.Table(tableName).Where("id=?", id).Updates(map[string]interface{}{"is_deleted": []byte{1}}).Error
	if err != nil {
		return
	}
	child := make([]int, 0)
	err = global.Eloquent.Table(tableName).Select("id").Where("pid=? AND is_deleted=?", id, 0).Find(&child).Error
	if err != nil || len(child) == 0 {
		return
	}
	for _, v := range child {
		err = DeleteChild(tableName, v)
		if err != nil {
			return
		}
	}
	return nil
}
