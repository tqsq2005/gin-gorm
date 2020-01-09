package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model

	TagId      int    `json:"tag_id" gorm:"index"`
	Tag		   Tag 	  `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	DeletedOn  int    `json:"deleted_on"`
	State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (article *Article) BeforeSave(scope *gorm.Scope) error  {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func GetArticles(pageNum int, limit int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(limit).Find(&articles)
	return
}

func GetArticleTotal(maps interface{}) (total int) {
	db.Model(&Article{}).Where(maps).Count(&total)
	return
}

func ExistArticleByTitle(title string) bool {
	var article Article
	db.Select("id").Where("title=?", title).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func ExistArticleById(id int) bool {
	var article Article
	db.First(&article, id)
	if article.ID > 0 {
		return true
	}
	return false
}

func GetArticle(id int) (article Article) {
	db.Where("id=?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

func EditArticle(id int, data interface{}) (article Article) {
	db.Model(&article).Where("id=?", id).Update(data)
	db.Model(&article).Related(&article.Tag)
	return
}

func AddArticle(data map[string]interface{}) (article Article) {
	article = Article{
		TagId: data["tag_id"].(int),
		Title: data["title"].(string),
		Desc:data["desc"].(string),
		Content: data["content"].(string),
		CreatedBy:data["created_by"].(string),
		State: data["state"].(int),
	}
	db.Create(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

func DeleteArticle(id int) bool {
	db.Where("id=?", id).Delete(&Article{})

	return true
}