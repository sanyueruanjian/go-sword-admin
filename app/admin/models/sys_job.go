package models

import (
	"project/app/admin/models/dto"
	"project/common/cache"
	orm "project/common/global"
	"project/utils"
)

// SysJob 岗位实体
type SysJob struct {
	*BaseModel
	Name     string `json:"name"`      //岗位名称
	Enabled  []byte `json:"enabled"`   //状态：1启用（默认）、0禁用
	JobSort  int    `json:"jobSort"`  //排序
	CreateBy int    `json:"create_by"` //创建者id
	UpdateBy int    `json:"update_by"` //更新者id
}

// SysJob 岗位表名
func (e *SysJob) TableName() string {
	return `sys_job`
}

// JobListDownload 导出岗位数据
func (e *SysJob) JobListDownload(p *dto.GetJobList) (jobList []*SysJob, err error) {
	orderJson, err := utils.OrderJson(p.Orders)
	if err != nil {
		return
	}
	orderRule := utils.GetOrderRule(orderJson)

	name := "%" + p.Name + "%"

	table := orm.Eloquent.Table(e.TableName()).Where("is_deleted=? AND name like ?", []byte{0}, name)
	if p.EndTime != 0 && p.StartTime != 0 {
		err = table.Where("create_time > ? AND create_time < ?", p.StartTime, p.EndTime).
			Order(orderRule).Find(&jobList).Error
	} else {
		err = table.Order(orderRule).Find(&jobList).Error
	}
	return
}

// GetJobList 查询岗位列表数据持久层
func (e *SysJob) GetJobList(p *dto.GetJobList, orderRule string) (jobList []*SysJob, count int64, err error) {
	name := "%" + p.Name + "%"
	table := orm.Eloquent.Table(e.TableName()).Where("is_deleted=? AND name like ?", []byte{0}, name)
	if p.EndTime != 0 && p.StartTime != 0 {
		err = table.Where("create_time > ? AND create_time < ?", p.StartTime, p.EndTime).Count(&count).
			Offset(p.Current - 1*p.Size).Limit(p.Size).Order(orderRule).Find(&jobList).Error
	} else {
		err = table.Count(&count).Offset((p.Current - 1) * p.Size).Limit(p.Size).Order(orderRule).Find(&jobList).Error
	}
	return
}

// GetJobList 查询岗位列表数据持久层
func (e *SysJob) GetJobEnabledList(p *dto.GetJobList, orderRule string) (jobList []*SysJob, count int64, err error) {
	err = orm.Eloquent.Table(e.TableName()).Where("is_deleted=?", []byte{0}).
	Count(&count).Order(orderRule).Find(&jobList).Error
	return
}

// DelJobById 删除岗位数据持久层
func (e *SysJob) DelJobById(userId int, ids *[]int) (err error) {
	// 查询修改岗位对应的用户id
	userIds, err := GetUserIdByJobId(ids)
	if err != nil {
		return
	}
	tx := orm.Eloquent.Begin()
	// 修改数据
	err = tx.Table(e.TableName()).Where("id in (?) AND is_deleted = ?", *ids, []byte{0}).Updates(map[string]interface{}{"is_deleted": 1, "update_by": userId}).Error
	if err != nil {
		return
	}
	// 删除岗位相关用户缓存
	err = cache.DelUserCacheById(cache.KeyUserJob, userIds)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// AddJob 新增岗位数据持久层
func (e *SysJob) AddJob () (err error) {
	err = orm.Eloquent.Create(e).Error
	return
}

// UpdateJob 更新岗位数据持久层
func (e *SysJob) UpdateJob (id int) (err error) {
	// 查询修改岗位对应的用户id
	userIds, err := GetUserIdByJobId(&[]int{id})
	if err != nil {
		return
	}
	tx := orm.Eloquent.Begin()
	// 修改数据
	err = tx.Table(e.TableName()).Where("id=? AND is_deleted=?", id, []byte{0}).Updates(e).Error
	if err != nil {
		return
	}
	// 删除岗位相关用户缓存
	err = cache.DelUserCacheById(cache.KeyUserJob, userIds)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// 查询岗位对应的用户id
func GetUserIdByJobId(ids *[]int) (userIds *[]int, err error) {
	userIds = new([]int)
	err = orm.Eloquent.Table("sys_users_jobs").Select("user_id").Where("job_id in (?)", *ids).Find(userIds).Error
	return
}
