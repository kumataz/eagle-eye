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



