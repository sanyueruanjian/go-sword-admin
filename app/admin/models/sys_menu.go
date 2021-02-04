package models

import (
	"encoding/json"
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/common/global"
	"project/utils"
)

const ForeNeed string = "ForeNeedMenu"

type SysMenu struct {
	*BaseModel
	Pid        int    `json:"pid"`        //上级菜单ID
	SubCount   int    `json:"sub_count"`  //子菜单数目
	Type       int    `json:"type"`       //菜单类型
	Title      string `json:"title"`      //菜单标题
	Name       string `json:"name"`       //组件名称
	Component  string `json:"component"`  //组件
	MenuSort   int    `json:"menu_sort"`  //排序
	Icon       string `json:"icon"`       //图标
	Path       string `json:"path"`       //链接地址
	IFrame     []byte `json:"i_frame"`    //是否外链
	Cache      []byte `json:"cache"`      //缓存
	Hidden     []byte `json:"hidden"`     //隐藏
	Permission string `json:"permission"` //权限
	CreateBy   int    `json:"create_by"`  //
	UpdateBy   int    `json:"update_by"`  //
}

type RedisForeNeedMenu struct {
	SelectForeNeedMenuList []*bo.SelectForeNeedMenuBo
}

func (m *SysMenu) TableName() string {
	return "sys_menu"
}

func (m *SysMenu) InsertMenu() error {
	err := global.Eloquent.Create(&m).Error
	if err != nil {
		return err
	}
	if err := global.Rdb.Del(ForeNeed).Err(); err != nil {
		return err
	}
	return nil
}

func (m *SysMenu) SelectMenu(p *dto.SelectMenuDto) (data []*SysMenu, err error) {
	//排序条件
	var orderJson []bo.Order
	orderJson, err = utils.OrderJson(p.Orders)
	orderRule := utils.GetOrderRule(orderJson)
	//模糊条件
	blurry := "%" + p.Blurry + "%"
	//时间条件
	if p.EndTime != 0 && p.StatTime != 0 {
		if err := global.Eloquent.Table("sys_menu").Where("is_deleted=? AND create_time > ? AND create_time < ? AND title like ?", []byte{0}, p.StatTime, p.EndTime, blurry).
			Limit(p.Size).Offset(p.Current - 1*p.Size).Order(orderRule).Find(&data).Error; err != nil {
			return nil, err
		}
	} else {
		if err := global.Eloquent.Table("sys_menu").Where("is_deleted=? AND title like ?", []byte{0}, blurry).
			Limit(p.Size).Offset(p.Current - 1*p.Size).Order(orderRule).Find(&data).Error; err != nil {
			return nil, err
		}
	}
	return data, err
}

//删除菜单
func (m *SysMenu) DeleteMenu(ids []int) (err error) {
	for _, v := range ids {
		err = utils.DeleteChild(m.TableName(), v)
		if err != nil {
			return err
		}
	}
	if err := global.Rdb.Del(ForeNeed).Err(); err != nil {
		return err
	}
	return nil
}

//更新菜单
func (m *SysMenu) UpdateMenu(p *dto.UpdateMenuDto, userId int) (err error) {
	err = global.Eloquent.Table("sys_menu").Where("id=?", p.ID).Updates(map[string]interface{}{
		"pid":        p.Pid,
		"sub_count":  p.SubCount,
		"type":       p.Type,
		"title":      p.Title,
		"name":       p.Name,
		"component":  p.Component,
		"menu_sort":  p.MenuSort,
		"icon":       p.Icon,
		"permission": p.Permission,
		"path":       p.Path,
		"update_by":  userId,
		"i_frame":    utils.BoolIntoByte(p.Iframe),
		"cache":      utils.BoolIntoByte(p.Cache),
		"hidden":     utils.BoolIntoByte(p.Iframe),
	}).Error
	if err := global.Rdb.Del(ForeNeed).Err(); err != nil {
		return err
	}
	return nil
}

//查找前端所需菜单
func (m *SysMenu) SelectForeNeedMenu(user *RedisUserInfo) (data []*bo.SelectForeNeedMenuBo, err error) {
	////检查缓存有无,有的话从缓存中读取
	var val []byte
	if global.Rdb.Exists(ForeNeed).Val() == 1 {
		val, err = global.Rdb.Get(ForeNeed).Bytes()
		redisForeNeedMenu := new(RedisForeNeedMenu)
		err = json.Unmarshal(val, redisForeNeedMenu)
		if err != nil {
			return nil, err
		}
		data = redisForeNeedMenu.SelectForeNeedMenuList
		if data != nil {
			return data, nil
		}
	}
	//查找角色Id
	parentIds := make([]int, 0)
	err = global.Eloquent.Table("sys_roles_menus").Select("menu_id").Where("role_id=?", 1).Joins("left join sys_menu "+
		"on sys_menu.id=sys_roles_menus.menu_id").Where("sys_menu.type=? AND sys_menu.is_deleted=?", 0, []byte{0}).Find(&parentIds).Error
	if err != nil {
		return
	}
	var results []*bo.SelectForeNeedMenuBo
	for _, pid := range parentIds {
		parentMenu := new(SysMenu)
		//查找父亲信息
		childMenus := make([]*SysMenu, 0)
		childS := make([]*bo.Children, 0)
		err = global.Eloquent.Table("sys_menu").Where("id=? AND is_deleted=?", pid, []byte{0}).First(parentMenu).Error
		if err != nil {
			return
		}
		result := &bo.SelectForeNeedMenuBo{
			Hidden:    utils.ByteIntoBool(parentMenu.Hidden),
			Component: parentMenu.Component,
			Name:      parentMenu.Name,
			Path:      "/" + parentMenu.Path,
			Meta: &bo.Meta{
				Icon:    parentMenu.Icon,
				NoCache: !utils.ByteIntoBool(parentMenu.Cache),
				Title:   parentMenu.Title,
			},
		}
		//查找多个父级菜单的子信息
		err = global.Eloquent.Table("sys_menu").Where("pid=? AND is_deleted=?", pid, []byte{0}).Find(&childMenus).Error
		if err != nil {
			return
		}
		result.AlwaysShow = false
		result.Redirect = "noredict"
		if len(childMenus) != 0 {
			result.AlwaysShow = true
			result.Redirect = "index"
		}
		//查询下一级
		for _, menu := range childMenus {
			child := &bo.Children{
				Component: menu.Component,
				Hidden:    utils.ByteIntoBool(menu.Hidden),
				Meta: &bo.Meta{
					Icon:    menu.Icon,
					NoCache: !utils.ByteIntoBool(menu.Cache),
					Title:   menu.Title,
				},
			}
			childS = append(childS, child)
		}
		result.Children = childS
		results = append(results, result)
	}
	//序列化结果
	var foreNeedMenu []byte
	redisForeNeedMenu := new(RedisForeNeedMenu)
	redisForeNeedMenu.SelectForeNeedMenuList = results
	foreNeedMenu, err = json.Marshal(redisForeNeedMenu)
	if err != nil {
		return nil, err
	}
	//添加缓存
	err = global.Rdb.Set(ForeNeed, foreNeedMenu, 0).Err()
	if err != nil {
		return nil, err
	}
	return results, nil
}
