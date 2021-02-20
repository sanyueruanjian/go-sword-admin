package service

import (
	"errors"
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
func (e *Job) GetJobList(p *dto.GetJobList) (*bo.GetJob, error) {
	job := new(models.SysJob)
	orderJson, err := utils.OrderJson(p.Orders)
	if err != nil {
		return nil, err
	}
	orderRule := utils.GetOrderRule(orderJson)
	res := new(bo.GetJob)
	if p.Size==9999 && p.Current==0 && p.Page==0 {
		jobList, count, err := job.GetJobEnabledList(p, orderRule)
		if err != nil {
			return nil, err
		}
		var getJobList []*bo.GetJobList
		for _, job := range jobList {
			getJobList = append(getJobList, &bo.GetJobList{
				Id:         job.ID,
				JobSort:    job.JobSort,
				CreateBy:   job.CreateBy,
				UpdateBy:   job.UpdateBy,
				CreateTime: job.CreateTime,
				UpdateTime: job.UpdateTime,
				Enabled:    utils.ByteIntoBool(job.Enabled),
				Name:       job.Name,
			})
		}

		res.Size = p.Size
		res.Current = p.Current
		res.Orders = orderJson
		res.Records = getJobList
		res.Total = int(count)
		res.Pages = p.Page

	} else if p.Size==0 || p.Current==0 {
		return nil, errors.New("参数缺失")
	} else {
		jobList, count, err := job.GetJobList(p, orderRule)
		if err != nil {
			return nil, err
		}
		var getJobList []*bo.GetJobList
		for _, job := range jobList {
			getJobList = append(getJobList, &bo.GetJobList{
				Id:         job.ID,
				JobSort:    job.JobSort,
				CreateBy:   job.CreateBy,
				UpdateBy:   job.UpdateBy,
				CreateTime: job.CreateTime,
				UpdateTime: job.UpdateTime,
				Enabled:    utils.ByteIntoBool(job.Enabled),
				Name:       job.Name,
			})
		}

		res.Size = p.Size
		res.Current = p.Current
		res.Orders = orderJson
		res.Records = getJobList
		res.Total = int(count)
		res.Pages = utils.PagesCount(res.Total, p.Size)
	}
	return res, nil
}

// DelJobById 删除岗位业务逻辑
func (e *Job) DelJobById(userId int, ids *[]int) (count *int64, err error) {
	count = new(int64)
	job := new(models.SysJob)
	err = job.QueryRelationshipJob(count, ids)
	if err != nil || *count > 0 {
		return
	}
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
