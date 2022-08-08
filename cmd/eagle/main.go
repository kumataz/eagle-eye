/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-06-01 20:15:04
 */
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/kumataahh/eagle-eye/internal/router"
	_ "github.com/kumataahh/eagle-eye/pkg/cron"

	"github.com/kumataahh/eagle-eye/internal"

	"github.com/gin-gonic/gin"
)

var (
	release bool = true
)

func main() {
	showLogo()

	//判断是否编译线上版本
	if release {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	app := internal.Application{
		Route: router.Init(),
	}
	app.Run()

}

func showLogo() {
	fmt.Println("  EAGELE - EYE !!   ")
}
