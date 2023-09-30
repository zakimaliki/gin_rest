package models

import (
	"gin_golang/src/config"

	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title string
	Slug  string `gorm:"unique_index"`
	Desc  string `gorm:"type:text"`
}

func SelectAll() *gorm.DB {
	items := []Article{}
	return config.DB.Find(&items)
}

func Select(slug string) *gorm.DB {
	var item Article
	// var item models.Article
	// if config.DB.First(&item, "slug = ?", slug).RecordNotFound() {
	// 	c.JSON(404, gin.H{
	// 		"status":  "Error",
	// 		"message": "record not found",
	// 	})
	// 	c.Abort()
	// 	return
	// }
	return config.DB.First(&item, "slug = ?", slug)
}

func Post(item *Article) *gorm.DB {
	return config.DB.Create(&item)
}

func Updates(id string, newArticle *Article) *gorm.DB {
	var item Article
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newArticle)
}

func Deletes(id string) *gorm.DB {
	var item Article
	return config.DB.Delete(&item, "id = ?", id)
}

func FindData(title string) *gorm.DB {
	items := []Article{}
	title = "%" + title + "%"
	return config.DB.Where("title LIKE ?", title).Find(&items)
}

func FindCond(sort string, limit int, offset int) *gorm.DB {
	items := []Article{}
	return config.DB.Order(sort).Limit(limit).Offset(offset).Find(&items)

}

func CountData() int {
	var result int
	config.DB.Table("articles").Count(&result)
	return result

}
