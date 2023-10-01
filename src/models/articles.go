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
	// items := []Article{}
	// config.DB.Raw("SELECT * FROM articles").Scan(&items)
	// return items

	items := []Article{}
	return config.DB.Find(&items)
}

func Select(slug string) *gorm.DB {
	// items := []Article{}
	// config.DB.Raw("SELECT * FROM articles WHERE slug = ?", slug).Scan(&items)
	// return items
	var item Article
	return config.DB.First(&item, "slug = ?", slug)
}

func Post(item *Article) *gorm.DB {
	// items := []Article{}
	// config.DB.Raw("INSERT INTO `articles` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `slug`, `description`) VALUES(Null, NULL, NULL, NULL, ?, ?, ?)", Title, Slug, Desc).Scan(&items)
	// return items
	return config.DB.Create(&item)
}

func Updates(id string, newArticle *Article) *gorm.DB {
	// items := []Article{}
	// config.DB.Raw("UPDATE articles SET title = ?, slug = ? , description = ? WHERE id = ?", Title, Slug, Desc, Id).Scan(&items)
	// return items
	var item Article
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newArticle)
}

func Deletes(id string) *gorm.DB {
	// items := []Article{}
	// config.DB.Raw("DELETE FROM articles WHERE id = ?", Id).Scan(&items)
	// return items
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
