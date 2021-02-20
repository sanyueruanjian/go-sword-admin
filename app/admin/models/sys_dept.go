package models

import (
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/common/global"
	orm "project/common/global"
	"project/utils"

	"gorm.io/gorm"
)

type SysDept struct {
	BaseModel
	Name     string `json:"name"`                              //名称
	Pid      int    `json:"pid"`                               //上级部门（顶级部门为0，默认为0）
	SubCount int    `json:"sub_count" gorm:"default:0"`        //子部门数目
	DeptSort int    `json:"deptSort"`                          //排序
	CreateBy int    `json:"create_by"`                         //创建者
	UpdateBy int    `json:"update_by"`                         //更新者
	Enabled  []byte `json:"enabled"  gorm:"default:[]byte{0}"` //状态：1启用（默认）、0禁用
}

// SysDept 部门表名
func (d *SysDept) TableName() string {
	return `sys_dept`
}

// 查询部门根据部门id
func (d *SysDept) GetDeptById() (err error) {
	err = orm.Eloquent.Table(d.TableName()).Where("id = ? AND is_deleted = 0", d.ID, 0).Find(d).Error
	return
}

// 查询与该部门相关的userId
func (d *SysDept) GetDeptUserListById() (ids []int, err error) {
	err = orm.Eloquent.Table("sys_name").Select("id").Where("dept_id = ? AND is_deleted = 0", d.ID, 0).Find(ids).Error
	return
}

// 查询部门
func (d *SysDept) SelectDeptListByPid(de *dto.SelectDeptDto, orderJson []bo.Order) (sysDeptList *[]SysDept, count int64, err error) {
	sysDeptList = new([]SysDept) // 实例化

	// 排序规则
	order := utils.GetOrderRule(orderJson)

	// 查询pid下的子部门数据
	if de.Pid >= 0 {
		//数据库查询
		err = orm.Eloquent.Table(d.TableName()).Where("pid = ? AND is_deleted = ?", de.Pid, 0).Count(&count).
			Order(order).Limit(de.Size).Offset((de.Current - 1) * de.Size).Find(sysDeptList).Error
		return
	}

	return
}

// 模糊查询和时间
func (d *SysDept) SelectDeptListByNameTime(de *dto.SelectDeptDto, orderJson []bo.Order) (sysDeptList *[]SysDept, count int64, err error) {
	sysDeptList = new([]SysDept) // 实例化

	// 排序规则
	order := utils.GetOrderRule(orderJson)

	// 模糊查询
	blurry := "%" + de.Name + "%"

	// 时间条件
	if de.EndTime != 0 && de.StartTime != 0 {
		err = global.Eloquent.Table(d.TableName()).Where("pid = ? AND is_deleted=? AND create_time > ? AND create_time < ? AND name like ?", 0, 0, de.StartTime, de.EndTime, blurry).
			Order(order).Count(&count).Limit(de.Size).Offset((de.Current - 1) * de.Size).Find(sysDeptList).Error
		return
	} else {
		// 数据库查询
		err = global.Eloquent.Table(d.TableName()).Where("pid = ? AND is_deleted=? AND name like ?", 0, 0, blurry).
			Order(order).Count(&count).Limit(de.Size).Offset((de.Current - 1) * de.Size).Find(sysDeptList).Error
	}
	return
}

// 新增部门
func (d *SysDept) InsertDept() (err error) {
	tx := global.Eloquent.Begin()
	err = tx.Table(d.TableName()).Create(d).Error
	if err != nil {
		tx.Rollback()
		return
	}
	if d.Pid != 0 {
		err = tx.Table(d.TableName()).Where("id = ?", d.Pid).Update("sub_count", gorm.Expr("sub_count + ?", 1)).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}
	err = tx.Commit().Error
	return
}

// 修改部门
func (d *SysDept) UpdateDept(de *dto.UpdateDeptDto) (err error) {
	err = global.Eloquent.Table(d.TableName()).Where("id=? AND is_deleted=?", de.ID, 0).Updates(&SysDept{
		Pid:      *de.Pid,
		Name:     de.Name,
		DeptSort: *de.DeptSort,
		UpdateBy: de.UpdateBy,
		Enabled:  utils.StrBoolIntoByte(de.Enabled),
	}).Error
	return
}

// 删除部门
func (d *SysDept) DeleteDept(ids *[]int, userId int) (count int64, err error) {
	child := new([]int)
	err = global.Eloquent.Table(d.TableName()).Select("id").Where("pid IN (?) AND is_deleted = ?", *ids, 0).Find(child).Error
	if err != nil {
		return
	}
	*child = append(*child, *ids...)
	err = global.Eloquent.Table("sys_user").Where("dept_id IN (?) AND is_deleted = ?", child, 0).Count(&count).Error
	if err != nil || count > 0 {
		return
	}

	tx := global.Eloquent.Begin()
	err = tx.Table(d.TableName()).Where("id IN (?)", *ids).Updates(&SysDept{
		UpdateBy: userId,
		BaseModel: BaseModel{
			IsDeleted: []byte{1},
		},
	}).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Table(d.TableName()).Where("pid IN (?)", *ids).Update("sub_count", gorm.Expr("sub_count - ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	return
}

func (d *SysDept) SuperiorDept(ids *[]int) (sysDeptList *[]SysDept, err error) {
	sysDeptList = new([]SysDept) // 实例化

	err = global.Eloquent.Table(d.TableName()).Where("pid = ? AND is_deleted = ?", 0, 0).
		Order("dept_sort").Find(sysDeptList).Error
	return
}

// 获取要下载的数据
func (d *SysDept) DownloadDept(de *dto.SelectDeptDto, orderJson []bo.Order) (sysDeptList []*SysDept, err error) {
	order := utils.GetOrderRule(orderJson)

	// 模糊查询
	blurry := "%" + de.Name + "%"

	// 时间条件
	if de.EndTime != 0 && de.StartTime != 0 {
		err = global.Eloquent.Table(d.TableName()).Where("pid = ? AND is_deleted=? AND create_time > ? AND create_time < ? AND title like ?", 0, 0, de.StartTime, de.EndTime, blurry).
			Order(order).Limit(de.Size).Offset((de.Current - 1) * de.Size).Find(&sysDeptList).Error
		return
	} else {
		err = global.Eloquent.Table(d.TableName()).Where("pid = ? AND is_deleted=? AND name like ?", 0, 0, blurry).
			Order(order).Limit(de.Size).Offset((de.Current - 1) * de.Size).Find(&sysDeptList).Error
	}
	return
}
