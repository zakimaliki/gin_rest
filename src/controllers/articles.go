package controllers

import (
	"fmt"
	"gin_golang/src/models"
	"log"
	"net/http"
	"strings"

	// "path/filepath"

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

func Uploadfile(c *gin.Context) {

	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	path := "src/uploads" + file.Filename

	c.SaveUploadedFile(file, path)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func FindArticle(c *gin.Context) {
	keyword := c.Query("search")
	res := models.FindData(keyword)

	c.JSON(202, gin.H{
		"status": "Berhasil",
		"data":   res,
	})
}

func FindTest(c *gin.Context) {
	sort := c.Query("sort")
	sortby := c.Query("sortby")
	sort = sort + " " + strings.ToLower(sortby)
	res := models.FindCond(sort)

	c.JSON(202, gin.H{
		"status": "Berhasil",
		"data":   res,
	})
}
