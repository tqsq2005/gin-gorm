package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	
	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeSave(scope *gorm.Scope) error  {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag)  {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int)  {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

//判断标签中的名称是否已存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

//根据主键获取记录
func ExistTagByID(id int) bool {
	var tag Tag
	db.First(&tag, id)
	if tag.ID > 0 {
		return true
	}
	return false
}

//创建标签
func AddTag(name string, state int, createdBy string) Tag {
	tag := Tag{
		Name:name,
		State:state,
		CreatedBy:createdBy,
	}
	db.Create(&tag)
	return tag
}

//修改标签
func EditTag(id int, data map[string]interface{}) Tag {
	var tag Tag
	db.Model(&tag).Where("id=?", id).Updates(data)
	return tag
}

//删除标签
func DeleteTag(id int) bool {
	var tag Tag
	db.Where("id=?", id).Delete(&tag)
	return true
}