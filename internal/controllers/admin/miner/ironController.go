/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-10-17 14:18:10
 */
package miner

import (
	// "fmt"
	"net/http"
	// "strconv"

	"github.com/kumataahh/eagle-eye/internal/controllers/admin"
	"github.com/kumataahh/eagle-eye/internal/models"
	services "github.com/kumataahh/eagle-eye/internal/services/admin"
	"github.com/kumataahh/eagle-eye/pkg/paginater"

	"github.com/gin-gonic/gin"
)

type ironController struct {
	admin.BaseController
}

var Iron = ironController{}

func (con ironController) List(c *gin.Context) {
	var (
		err         error
		req         models.MinerIndexReq
		minerList []models.Collections
	)

	err = con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	adminDb := services.MinerService.GetMiners(req)

	minerData := paginater.PageOperation(c, adminDb, 50, &minerList)

	c.HTML(http.StatusOK, "miner/miner.html", gin.H{
		"minerData": minerData,
		"created_at":  c.Query("created_at"),
		"title":       c.Query("title"),
	})
}


// func (con ironController) Add(c *gin.Context) {
// 	var article models.Article
// 	c.HTML(http.StatusOK, "article/article_form.html", gin.H{
// 		"article": article,
// 	})
// }

// func (con ironController) Edit(c *gin.Context) {
// 	articel_id := c.Query("article_id")

// 	id, _ := strconv.Atoi(articel_id)

// 	article, err := services.ArticleService.GetArticle(uint(id))
// 	if err != nil {
// 		con.Error(c, err.Error())
// 	}

// 	c.HTML(http.StatusOK, "article/article_form.html", gin.H{
// 		"article": article,
// 		"url":     c.Request.RequestURI,
// 	})
// }


// func (con ironController) Save(c *gin.Context) {
// 	var (
// 		req models.ArticleReq
// 		err error
// 	)

// 	con.FormBind(c, &req)
// 	err = services.ArticleService.SaveArticle(req)
// 	if err != nil {
// 		con.Error(c, err.Error())
// 	}

// 	con.Success(c, "/admin/article/list", "添加成功")
// }

// func (con ironController) Del(c *gin.Context) {
// 	id := c.Query("article_id")
// 	fmt.Println(id)
// 	articleId, _ := strconv.Atoi(id)

// 	err := services.ArticleService.DelArticle(articleId)
// 	if err != nil {
// 		con.Error(c, "删除失败")
// 	} else {
// 		con.Success(c, "", "删除成功")
// 	}
// }
