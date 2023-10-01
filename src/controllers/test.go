package controllers

import (
	"gin_golang/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func TestList(c *gin.Context) {
	res := models.RawList()

	c.JSON(202, gin.H{
		"status": "Berhasil",
		"data":   res,
	})
}

func Test(c *gin.Context) {
	slug := c.Param("slug")
	res := models.Raw(slug)

	c.JSON(202, gin.H{
		"status": "Berhasil",
		"data":   res,
	})
}

func TestPost(c *gin.Context) {
	Title := c.PostForm("title")
	Slug := slug.Make(c.PostForm("title"))
	Desc := c.PostForm("description")

	res := models.RawPost(Title, Slug, Desc)

	c.JSON(201, gin.H{
		"status": "Berhasil",
		"data":   res,
	})
}

func TestUpdate(c *gin.Context) {
	IdOld := c.Param("id")
	Id, _ := strconv.Atoi(IdOld)
	Title := c.PostForm("title")
	Slug := slug.Make(c.PostForm("title"))
	Desc := c.PostForm("description")

	res := models.RawPut(Title, Slug, Desc, Id)

	c.JSON(202, gin.H{
		"status": "Berhasil",
		"data":   res,
	})
}

func TestDelete(c *gin.Context) {
	IdOld := c.Param("id")
	Id, _ := strconv.Atoi(IdOld)
	res := models.RawDelete(Id)

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   res,
	})
}
