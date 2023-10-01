package models

import (
	"gin_golang/src/config"
)

func RawList() []Article {
	items := []Article{}
	config.DB.Raw("SELECT * FROM articles").Scan(&items)
	return items
}

func Raw(slug string) []Article {
	items := []Article{}
	config.DB.Raw("SELECT * FROM articles WHERE slug = ?", slug).Scan(&items)
	return items
}

func RawPost(Title string, Slug string, Desc string) []Article {
	items := []Article{}
	config.DB.Raw("INSERT INTO `articles` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `slug`, `desc`) VALUES(Null, NULL, NULL, NULL, ?, ?, ?)", Title, Slug, Desc).Scan(&items)
	return items
}

func RawPut(Title string, Slug string, Desc string, Id int) []Article {
	items := []Article{}
	config.DB.Raw("UPDATE articles SET title = ?, slug = ? , desc = ? WHERE id = ?", Title, Slug, Desc, Id).Scan(&items)
	return items
}

func RawDelete(Id int) []Article {
	items := []Article{}
	config.DB.Raw("DELETE FROM articles WHERE id = ?", Id).Scan(&items)
	return items
}
