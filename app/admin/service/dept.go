package service

import (
	"io"
	"project/app/admin/models"
	"project/app/admin/models/bo"
	"project/app/admin/models/cache"
	"project/app/admin/models/dto"
	"project/utils"

	"go.uber.org/zap"
)

type Dept struct {
}

func (d Dept) SelectDeptList(de *dto.SelectDeptDto, orderData []bo.Order) (data *bo.SelectDeptListBo, err error) {
	// 声明所需变量，开辟空间
	data = new(bo.SelectDeptListBo)
	deptList := new([]bo.RecordDept)
	dept := new(models.SysDept)
	sysDeptList := new([]models.SysDept)
	var count int64

	// 数据查询 判断条件
	tag := de.Name != "" || de.StartTime != 0 || de.EndTime != 0

	// 进dao层
	if tag {
		// 模糊查询直接过数据库
		sysDeptList, count, err = dept.SelectDeptListByNameTime(de, orderData)
	} else {
		// 非模糊查询先过缓存
		*deptList, err = cache.GetRedisDeptByPid(de.Pid)
		if err == nil && len(*deptList) > 0 {
			// 封装paging
			data.Orders = orderData
			data.Current = de.Current
			data.Total = len(*deptList)
			data.Size = de.Size
			data.Pages = utils.PagesCount(data.Total, de.Size)
			data.Records = *deptList
			return
		}
		// 缓存有问题
		if err != nil {
			zap.L().Error("GetRedisDeptByPid failed", zap.Error(err))
			err = nil
		}
		_ = cache.DeleteRedisDeptByPid(de.Pid)
		sysDeptList, count, err = dept.SelectDeptListByPid(de, orderData)
	}

	// 数据库错误
	if err != nil {
		zap.L().Error("SelectDeptDao Select failed", zap.Error(err))
		return
	}

	// 封装bo数据传输对象
	if len(*sysDeptList) > 0 {
		deptList = modelToBo(sysDeptList)
		// 子部门缓存
		if !tag {
			_ = cache.DeleteRedisDeptByPid(de.Pid)

			err = cache.SetRedisDeptByPid(de.Pid, deptList)
			if err != nil {
				zap.L().Error("SetRedisDeptByPid failed", zap.Error(err))
				err = nil
			}

			err = cache.SetRedisDeptList(deptList)
			if err != nil {
				zap.L().Error("SetRedisDeptList failed", zap.Error(err))
				err = nil
			}
		}

		// 封装paging
		data.Orders = orderData
		data.Current = de.Current
		data.Total = int(count)
		data.Size = de.Size
		data.Pages = utils.PagesCount(data.Total, de.Size)
		data.Records = *deptList
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

	// 删除缓存
	err = cache.DeleteRedisDeptByPid(*de.Pid)
	if err != nil {
		return
	}
	// 存入数据库
	err = dept.InsertDept()
	return
}

// 修改部门
func (d Dept) UpdateDept(de *dto.UpdateDeptDto) (err error) {
	dept := new(models.SysDept)

	// 删除缓存
	err = cache.DeleteRedisDeptByPid(*de.Pid)
	if err != nil {
		return
	}
	err = cache.DeleteRedisDeptById(de.ID)
	if err != nil {
		return
	}
	// 持久层
	err = dept.UpdateDept(de)
	return
}

// 删除部门
func (d Dept) DeleteDept(ids *[]int) (count int64, err error) {
	dept := new(models.SysDept)
	count, err = dept.DeleteDept(ids)
	if count == 0 {
	}
	return
}

func (d Dept) SuperiorDept(ids *[]int) (deptList *[]bo.RecordDept, err error) {
	// 数据查询
	deptList = new([]bo.RecordDept)
	sysDeptList := new([]models.SysDept)
	dept := new(models.SysDept)
	// 非模糊查询先过缓存
	*deptList, err = cache.GetRedisDeptByPid(0)
	if err == nil && len(*deptList) > 0 {
		return
	}
	// 缓存有问题
	if err != nil {
		zap.L().Error("GetRedisDeptByPid failed", zap.Error(err))
		err = nil
	}
	_ = cache.DeleteRedisDeptByPid(0)
	sysDeptList, err = dept.SuperiorDept(ids)
	if err != nil {
		return
	}

	// 封装bo数据传输对象
	if len(*sysDeptList) > 0 {
		deptList = modelToBo(sysDeptList)
		// 子部门缓存
		_ = cache.DeleteRedisDeptByPid(0)

		err = cache.SetRedisDeptByPid(0, deptList)
		if err != nil {
			zap.L().Error("SetRedisDeptByPid failed", zap.Error(err))
			err = nil
		}

		err = cache.SetRedisDeptList(deptList)
		if err != nil {
			zap.L().Error("SetRedisDeptList failed", zap.Error(err))
			err = nil
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
		deptList = append(deptList, bo.DownloadDeptList{
			Name:       dept.Name,
			Enabled:    utils.ByteEnabledToString(dept.Enabled),
			CreateTime: utils.UnixTimeToString(dept.CreateTime),
		})
	}

	// 生成excel
	content = utils.ToExcel([]string{`部门名称`, `部门状态`, `创建日期`}, deptList)
	return
}

func modelToBo(sysDeptList *[]models.SysDept) (deptList *[]bo.RecordDept) {
	var r bo.RecordDept
	deptList = new([]bo.RecordDept)
	for _, value := range *sysDeptList {
		r.ID = value.ID
		r.Pid = value.Pid
		r.Name = value.Name
		r.Label = value.Name
		r.Enabled = utils.ByteIntoBool(value.Enabled)
		if value.SubCount > 0 {
			r.HasChildren = true
			r.Leaf = false
		} else {
			r.HasChildren = false
			r.Leaf = true
		}
		r.CreateTime = value.CreateTime
		r.CreateBy = value.CreateBy
		r.UpdateTime = value.UpdateTime
		r.UpdateBy = value.UpdateBy
		// append
		*deptList = append(*deptList, r)
	}
	return
}
