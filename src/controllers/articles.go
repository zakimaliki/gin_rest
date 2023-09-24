package controllers

import (
	"gin_golang/src/models"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func GetListArticle(c *gin.Context) {
	res := models.SelectAll()
	c.JSON(200, gin.H{
		"status": "Berhasil",
		"data":   res,
	})
}

func GetArticle(c *gin.Context) {
	slug := c.Param("slug")
	res := models.Select(slug)

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   res,
	})
}

func PostArticle(c *gin.Context) {

	item := models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  slug.Make(c.PostForm("title")),
	}

	res := models.Post(&item)

	c.JSON(201, gin.H{
		"status": "Berhasil",
		"data":   res,
	})
}

func UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	newArticle := models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  slug.Make(c.PostForm("title")),
	}

	res := models.Updates(id, &newArticle)

	c.JSON(202, gin.H{
		"status": "Berhasil",
		"data":   res,
	})
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	res := models.Deletes(id)

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   res,
	})
}
