/*
 * @Description:
 * @Author: gphper
 * @Date: 2022-03-15 20:14:09
 */
package dao

import (
	"github.com/kumataahh/eagle-eye/internal/models"

	"gorm.io/gorm"
)

type adminUserDao struct {
	DB *gorm.DB
}

var AuDao = adminUserDao{DB: models.Db}
