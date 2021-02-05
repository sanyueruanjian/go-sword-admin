package models

import (
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/common/global"
	orm "project/common/global"
	"project/utils"
)

type SysDept struct {
	*BaseModel
	Name     string `json:"name"`                              //名称
	Pid      int    `json:"pid"`                               //上级部门（顶级部门为0，默认为0）
	SubCount int    `json:"sub_count" gorm:"default:0"`        //子部门数目
	DeptSort int    `json:"deptSort"`                          //排序
	CreateBy int    `json:"create_by"`                         //创建者
	UpdateBy int    `json:"update_by"`                         //更新者
	Enabled  []byte `json:"enabled"  gorm:"default:[]byte{0}"` //状态：1启用（默认）、0禁用
}

// 查询部门
func (d SysDept) SelectDeptList(de *dto.SelectDeptDto, orderJson []bo.Order) (sysDeptList []*SysDept, err error) {
	// 排序规则
	order := utils.GetOrderRule(orderJson)

	// enabled
	if de.Enabled == true {
		err = orm.Eloquent.Table("sys_dept").Where("enabled = ? AND pid = ? AND is_deleted = ?", 1, 0, []byte{0}).
			Limit(de.Size).Offset((de.Current - 1) * de.Size).Order(order).Find(&sysDeptList).Error
		return
	}

	// 查询pid下的子部门数据
	if de.Pid > 0 {
		err = orm.Eloquent.Table("sys_dept").Where("enabled = ? AND pid = ? AND is_deleted = ?", 1, de.Pid, []byte{0}).
			Limit(de.Size).Offset((de.Current - 1) * de.Size).Order(order).Find(&sysDeptList).Error
		return
	}

	// 模糊查询
	blurry := "%" + de.Name + "%"

	// 时间条件
	if de.EndTime != 0 && de.StartTime != 0 {
		err = global.Eloquent.Table("sys_dept").Where("pid = ? AND is_deleted=? AND create_time > ? AND create_time < ? AND title like ?", 0, []byte{0}, de.StartTime, de.EndTime, blurry).
			Limit(de.Size).Offset(de.Current - 1*de.Size).Order(order).Find(&sysDeptList).Error
		return
	} else {
		err = global.Eloquent.Table("sys_dept").Where("pid = ? AND is_deleted=? AND name like ?", 0, []byte{0}, blurry).
			Limit(de.Size).Offset(de.Current - 1*de.Size).Order(order).Find(&sysDeptList).Error
	}
	return
}

// 新增部门
func (d SysDept) InsertDept() (err error) {
	err = global.Eloquent.Table("sys_dept").Create(&d).Error
	return
}

// 修改部门
func (d SysDept) UpdateDept(de *dto.UpdateDeptDto) (err error) {
	err = global.Eloquent.Table("sys_dept").Where("id=? AND is_deleted=?", de.ID, 0).Updates(map[string]interface{}{
		"pid":       de.Pid,
		"sub_count": de.SubCount,
		"name":      de.Name,
		"dept_sort": de.DeptSort,
		"update_by": de.UpdatedBy,
		"enabled":   utils.StrBoolIntoByte(de.Enabled),
	}).Error
	return
}

// 删除部门
func (d SysDept) DeleteDept(ids *[]int) (err error) {
	err = global.Eloquent.Table("sys_dept").Where("id IN (?)", *ids).Updates(map[string]interface{}{"is_deleted": []byte{1}}).Error
	return
}

func (d SysDept) SuperiorDept(ids *[]int) (sysDeptList []*SysDept, err error) {
	err = orm.Eloquent.Table("sys_dept").Where("enabled = ? AND pid = ? AND is_deleted = ?", 1, 0, []byte{0}).
		Order("dept_sort").Find(&sysDeptList).Error
	return
}

// 获取要下载的数据
func (d SysDept) DownloadDept(de *dto.SelectDeptDto, orderJson []bo.Order) (sysDeptList []*SysDept, err error) {
	order := utils.GetOrderRule(orderJson)

	// 模糊查询
	blurry := "%" + de.Name + "%"

	// 时间条件
	if de.EndTime != 0 && de.StartTime != 0 {
		err = global.Eloquent.Table("sys_dept").Where("pid = ? AND is_deleted=? AND create_time > ? AND create_time < ? AND title like ?", 0, []byte{0}, de.StartTime, de.EndTime, blurry).
			Limit(de.Size).Offset(de.Current - 1*de.Size).Order(order).Find(&sysDeptList).Error
		return
	} else {
		err = global.Eloquent.Table("sys_dept").Where("pid = ? AND is_deleted=? AND name like ?", 0, []byte{0}, blurry).
			Limit(de.Size).Offset(de.Current - 1*de.Size).Order(order).Find(&sysDeptList).Error
	}
	return
}
