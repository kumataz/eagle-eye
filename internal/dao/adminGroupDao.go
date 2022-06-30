package dao

import (
	"github.com/kumataahh/eagle-eye/internal/models"

	"gorm.io/gorm"
)

type adminGroupDao struct {
	DB *gorm.DB
}

var AgDao = adminGroupDao{DB: models.Db}
