package controllers

import (
	"gin_golang/src/config"
	"gin_golang/src/models"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func GetHome(c *gin.Context) {
	items := []models.Article{}
	config.DB.Find(&items)
	c.JSON(200, gin.H{
		"status": "Berhasil",
		"data":   items,
	})
}

func GetArticle(c *gin.Context) {
	slug := c.Param("slug")

	var item models.Article

	if config.DB.First(&item, "slug = ?", slug).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "Error",
			"message": "record not found",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   item,
	})
}

func PostArticle(c *gin.Context) {

	item := models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  slug.Make(c.PostForm("title")),
	}

	config.DB.Create(&item)

	c.JSON(201, gin.H{
		"status": "Berhasil",
		"data":   item,
	})
}

func UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	newArticle := models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  slug.Make(c.PostForm("title")),
	}

	var item models.Article

	config.DB.Model(&item).Where("id = ?", id).Updates(&newArticle)

	c.JSON(202, gin.H{
		"status": "Berhasil",
		"data":   item,
	})
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	var item models.Article
	config.DB.Delete(models.Article{}, "id = ?", id)

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   item,
	})
}
