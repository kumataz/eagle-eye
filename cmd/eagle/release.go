// +build !release

package main

import (
	_ "github.com/kumataahh/eagle-eye/docs"
	"github.com/kumataahh/eagle-eye/internal/router"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	release = false
	router.SwagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
}
