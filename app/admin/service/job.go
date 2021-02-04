package service

import (
	"io"

	"project/app/admin/models"
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/utils"
)

type Job struct {
}

// JobListDownload 下载岗位列表业务逻辑
func (e *Job) JobListDownload(p *dto.GetJobList) (content io.ReadSeeker, err error) {
	var res []interface{}
	job := new(models.SysJob)
	jobList, err := job.JobListDownload(p)
	if err != nil {
		return
	}

	for _, job := range jobList {
		res = append(res, &bo.JobListDownload{
			Name:       job.Name,
			Enabled:    utils.ByteEnabledToString(job.Enabled),
			CreateTime: utils.UnixToFormatTime(job.CreateTime),
		})
	}

	content = utils.ToExcel([]string{`岗位名称`, `岗位状态`, `创建日期`}, res)
	return
}

// GetJobList 查询岗位列表业务逻辑
func (e *Job) GetJobList(p *dto.GetJobList) (res []*bo.GetJobList, err error) {
	job := new(models.SysJob)
	jobList, err := job.GetJobList(p)
	if err != nil {
		return
	}

	for _, job := range jobList {
		res = append(res, &bo.GetJobList{
			Name:    job.Name,
			Enabled: utils.ByteIntoInt(job.Enabled),
			JobSort: job.JobSort,
		})
	}

	return
}

// DelJobById 删除岗位业务逻辑
func (e *Job) DelJobById(userId int, ids []int) (err error) {
	job := new(models.SysJob)
	err = job.DelJobById(userId, ids)
	return
}

// AddJob 新增岗位业务逻辑
func (e *Job) AddJob(userId int, p *dto.AddJob) (err error) {
	job := new(models.SysJob)
	job.JobSort = p.JobSort
	job.Name = p.Name
	job.Enabled = utils.BoolIntoByte(p.Enabled)
	job.CreateBy = userId
	job.UpdateBy = userId
	err = job.AddJob()
	return
}

// Update 更新岗位业务逻辑
func (e *Job) Update(userId int, p *dto.UpdateJob) (err error) {
	job := new(models.SysJob)
	job.JobSort = p.JobSort
	job.Name = p.Name
	job.Enabled = utils.BoolIntoByte(p.Enabled)
	job.UpdateBy = userId
	err = job.UpdateJob(p.ID)
	return
}
