/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-10-17 14:18:10
 */
package miner

import (
	// "fmt"
	"net/http"
	"strings"
	"strconv"
	"fmt"

	"github.com/kumataahh/eagle-eye/internal/controllers/admin"
	"github.com/kumataahh/eagle-eye/internal/models"
	"github.com/kumataahh/eagle-eye/internal/database"
	services "github.com/kumataahh/eagle-eye/internal/services/admin"
	"github.com/kumataahh/eagle-eye/pkg/paginater"

	"github.com/gin-gonic/gin"
)

type ironController struct {
	admin.BaseController
}

var Iron = ironController{}
var param models.Collections

func (con ironController) PostJson(c *gin.Context) {

	param = models.Collections{}
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(500, gin.H{"err": err})
		  return
	}
	fmt.Println(param)
	fmt.Printf("param-type: %T\n", param)
	// fmt.Printf("param-string-type: %T\n", string(param))
	database.PostIronMysql(param)
	c.JSON(200, gin.H{
		"results": param,
	})
}

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
	})
}

func (con ironController) Del(c *gin.Context) {
	// 处理参数
	id := c.Query("machine_id")
	res := strings.TrimSpace(id)
    arr := strings.Split(res, ",")

	for i:= 0;i<len(arr);i++{
		machineId, _ := strconv.Atoi(arr[i])
		// 调用方法
		err := services.MinerService.DelMiner(machineId)
		if err != nil {
			con.Error(c, "del fail")
		} else {
			con.Success(c, "", "del success")
		}
	}
}



// func (con ironController) FindTag(c *gin.Context) {
// 	var (
// 		err         error
// 		req         models.MinerIndexReq
// 		tagList    []models.Collections
// 	)

// 	err = con.FormBind(c, &req)
// 	if err != nil {
// 		con.Error(c, err.Error())
// 		return
// 	}

// 	tag := c.Query("machine_tag")
// 	adminDb := services.MinerService.FindMiner(tag, req)

// 	tagData := paginater.PageOperation(c, adminDb, 50, &tagList)
// 	fmt.Printf("tagData type: %T", tagData)
// 	c.HTML(http.StatusOK, "miner/miner.html", gin.H{
// 		"tagData": tagData,
// 	})
// }



