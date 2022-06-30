/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-14 13:39:32
 */
package dao

import (
	"github.com/kumataahh/eagle-eye/internal/models"

	"gorm.io/gorm"
)

type minerDao struct {
	DB *gorm.DB
}

var MinerDao = minerDao{DB: models.Db}
