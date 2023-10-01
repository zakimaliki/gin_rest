package controllers

import (
	"fmt"
	"gin_golang/src/models"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

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
		Desc:  c.PostForm("description"),
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
		Desc:  c.PostForm("description"),
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
	cekFile := strings.Split(file.Filename, ".")

	if (cekFile[1] == "png") || (cekFile[1] == "jpg") || (cekFile[1] == "jpg") {
		timestampOld := int(time.Now().Unix())
		addFile := strconv.Itoa(timestampOld) + "_" + file.Filename

		path := "src/uploads/" + addFile
		c.SaveUploadedFile(file, path)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	} else {
		c.String(http.StatusOK, fmt.Sprintf("File Picture format must PNG, JPG , or JPEG"))
	}
}

func FindArticle(c *gin.Context) {
	keyword := c.Query("search")
	res := models.FindData(keyword)

	c.JSON(202, gin.H{
		"status": "Berhasil",
		"data":   res,
	})
}

func PaginatSortArticle(c *gin.Context) {
	pageOld := c.Query("page")
	limitOld := c.Query("limit")
	page, _ := strconv.Atoi(pageOld)
	limit, _ := strconv.Atoi(limitOld)
	offset := (page - 1) * limit
	sort := c.Query("sort")
	sortby := c.Query("sortby")
	sort = sortby + " " + strings.ToLower(sort)
	res := models.FindCond(sort, limit, offset)
	totalData := models.CountData()
	totalPage := math.Ceil(float64(totalData) / float64(limit))

	c.JSON(202, gin.H{
		"status":      "Berhasil",
		"data":        res,
		"currentPage": page,
		"limit":       limit,
		"totalData":   totalData,
		"totalPage":   totalPage,
	})
}
