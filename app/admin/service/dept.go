package service

import (
	"io"
	"project/app/admin/models"
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/utils"
)

type Dept struct {
}

func (d Dept) SelectDeptList(de *dto.SelectDeptDto, orderData []bo.Order) (data []*bo.RecordDept, err error) {
	// 数据查询
	dept := new(models.SysDept)
	sysDeptList, err := dept.SelectDeptList(de, orderData)
	if err != nil {
		return
	}

	// 封装bo数据传输对象
	if len(sysDeptList) > 0 {
		for _, value := range sysDeptList {
			recordDept := new(bo.RecordDept)
			recordDept.ID = value.ID
			if value.Pid != 0 {
				recordDept.Leaf = false
			} else {
				recordDept.Leaf = true
			}
			recordDept.Pid = value.Pid
			recordDept.Name = value.Name
			recordDept.Label = value.Name
			recordDept.Enabled = utils.ByteIntoBool(value.Enabled)
			if value.SubCount > 0 {
				recordDept.HasChildren = true
			} else {
				recordDept.HasChildren = false
			}
			recordDept.CreateTime = value.CreateTime
			recordDept.CreateBy = value.CreateBy
			recordDept.UpdateTime = value.UpdateTime
			recordDept.UpdateBy = value.UpdateBy

			// append
			data = append(data, recordDept)
		}
	}
	return
}

// 新增部门
func (d Dept) InsertDept(de *dto.InsertDeptDto, userId int) (err error) {
	// 实例化
	dept := new(models.SysDept)
	dept.DeptSort = de.DeptSort
	dept.Enabled = utils.StrBoolIntoByte(de.Enabled)
	dept.Pid = *de.Pid
	dept.Name = de.Name
	dept.SubCount = *de.SubCount
	dept.CreateBy = userId
	dept.UpdateBy = userId

	// 存入数据库
	err = dept.InsertDept()
	return err
}

// 修改部门
func (d Dept) UpdateDept(de *dto.UpdateDeptDto) (err error) {
	dept := new(models.SysDept)
	// 持久层
	err = dept.UpdateDept(de)
	return
}

// 删除部门
func (d Dept) DeleteDept(ids *[]int) (err error) {
	dept := new(models.SysDept)
	err = dept.DeleteDept(ids)
	return
}

func (d Dept) SuperiorDept(ids *[]int) (data []*bo.RecordDept, err error) {
	// 数据查询
	dept := new(models.SysDept)
	sysDeptList, err := dept.SuperiorDept(ids)
	if err != nil {
		return
	}

	// 封装bo数据传输对象
	if len(sysDeptList) > 0 {
		for _, value := range sysDeptList {
			recordDept := new(bo.RecordDept)
			recordDept.ID = value.ID
			if value.Pid != 0 {
				recordDept.Leaf = false
			} else {
				recordDept.Leaf = true
			}
			recordDept.Pid = value.Pid
			recordDept.Name = value.Name
			recordDept.Label = value.Name
			recordDept.Enabled = utils.ByteIntoBool(value.Enabled)
			if value.SubCount > 0 {
				recordDept.HasChildren = true
			} else {
				recordDept.HasChildren = false
			}
			recordDept.CreateTime = value.CreateTime
			recordDept.CreateBy = value.CreateBy
			recordDept.UpdateTime = value.UpdateTime
			recordDept.UpdateBy = value.UpdateBy

			// append
			data = append(data, recordDept)
		}
	}
	return
}

func (d Dept) DownloadDeptList(dt *dto.SelectDeptDto, orderJson []bo.Order) (content io.ReadSeeker, err error) {
	var deptList []interface{}
	dept := new(models.SysDept)
	// 数据库查询数据
	sysDeptList, err := dept.DownloadDept(dt, orderJson)
	// 对数据库取出的sysDept封住存入deptList
	for _, dept := range sysDeptList {
		deptList = append(deptList, &bo.DownloadDeptList{
			Name:       dept.Name,
			Enabled:    utils.ByteEnabledToString(dept.Enabled),
			CreateTime: utils.UnixTimeToString(dept.CreateTime),
		})
	}

	// 生成excel
	content = utils.ToExcel([]string{`部门名称`, `部门状态`, `创建日期`}, deptList)
	return
}
