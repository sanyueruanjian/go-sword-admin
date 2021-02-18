package cache

import (
	"encoding/json"
	"project/app/admin/models/bo"
	"project/common/global"
	"strconv"
)

const (
	DeptIdKeyFore  = "dept::id:"
	DeptPidKeyFore = "dept::pid:"
)

// dept::pid:{number} 设置子部门缓存
func SetRedisDeptByPid(pid int, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	strId := strconv.Itoa(pid)
	return global.Rdb.Set(DeptPidKeyFore+strId, data, 0).Err()
}

// dept::id:{number} 批量设置部门缓存
func SetRedisDeptList(value *[]bo.RecordDept) (err error) {
	var id int
	pipLine := global.Rdb.TxPipeline()
	for _, v := range *value {
		id = v.ID
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		strId := strconv.Itoa(id)
		pipLine.Set(DeptIdKeyFore+strId, data, 0)
	}
	_, err = pipLine.Exec()
	return
}

// dept::pid:{number} 获取子部门缓存
func GetRedisDeptByPid(pid int) (value []bo.RecordDept, err error) {
	strId := strconv.Itoa(pid)
	data, err := global.Rdb.Get(DeptPidKeyFore + strId).Bytes()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &value)
	return
}

// dept::pid:{number} 删除子部门缓存
func DeleteRedisDeptByPid(pid int) error {
	var strId = strconv.Itoa(pid)
	return global.Rdb.Del(DeptPidKeyFore + strId).Err()
}

// dept::pid:{number} 删除子部门缓存
func DeleteRedisDeptByPids(pids []int) (err error) {
	for _, pid := range pids {
		var strId = strconv.Itoa(pid)
		err = global.Rdb.Del(DeptPidKeyFore + strId).Err()
		if err != nil {
			return
		}
	}
	return
}

// dept::id:{number} 设置某个部门缓存
func SetRedisDeptById(id int, value interface{}) {
	data, err := json.Marshal(value)
	if err != nil {
		return
	}
	strId := strconv.Itoa(id)
	global.Rdb.Set(DeptIdKeyFore+strId, data, 0)
}

// dept::id:{number} 获取某个部门缓存
func GetRedisDeptById(id int) (value bo.RecordDept, err error) {
	strId := strconv.Itoa(id)
	data, err := global.Rdb.Get(DeptIdKeyFore + strId).Bytes()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &value)
	return
}

// dept::id:{number} 获取某些部门缓存
func GetRedisDeptByIdList(ids []int) (value []bo.RecordDept, err error) {
	for id := range ids {
		var dept bo.RecordDept
		strId := strconv.Itoa(id)
		data, err := global.Rdb.Get(DeptIdKeyFore + strId).Bytes()
		if err != nil {
			return value, err
		}
		err = json.Unmarshal(data, &dept)
		if err != nil {
			return value, err
		}
		value = append(value, dept)
	}
	return
}

// dept::pid:{number} 删除某个部门缓存
func DeleteRedisDeptById(id int) (err error) {
	var strId = strconv.Itoa(id)
	return global.Rdb.Del(DeptIdKeyFore + strId).Err()
}

// dept::pid:{number} 删除某些部门缓存
func DeleteRedisDeptByIds(ids []int) (err error) {
	for _, id := range ids {
		var strId = strconv.Itoa(id)
		err = global.Rdb.Del(DeptIdKeyFore + strId).Err()
		if err != nil {
			return
		}
	}
	return
}
