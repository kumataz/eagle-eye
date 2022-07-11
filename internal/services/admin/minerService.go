/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-14 13:29:28
 */
package admin

import (
	// "strings"
	// "fmt"
	"github.com/kumataahh/eagle-eye/internal/dao"
	"github.com/kumataahh/eagle-eye/internal/models"

	"gorm.io/gorm"
)

type minerService struct{}

var MinerService = minerService{}

// func (ser *minerService) GetMiner(minerId uint) (miner models.Miner, err error) {
// 	miner.MinerId = minerId
// 	err = dao.MinerDao.DB.First(&miner).Error
// 	if err != nil {
// 		return models.Miner{}, err
// 	}
// 	return
// }

func (ser *minerService) GetMiners(req models.MinerIndexReq) (db *gorm.DB) {
	// req not use here
	// db = dao.MinerDao.DB.Table("miner_ironfish")
	var collect models.Collections
	// id升序再按status把红色推上去
	db = dao.MinerDao.DB.Order("work_status desc").Order("machine_id asc").Find(&collect)
	// db = dao.MinerDao.DB.Where("machine_tag = ?", "miner").Find(&collect)

	return
}

func (ser *minerService) DelMiner(id int) (err error) {
	var collect models.Collections
	collect.MachineId = int(id)
	err = dao.MinerDao.DB.Where("machine_id = ?", id).Delete(&collect).Error
	// if err != nil {
	// 	return models.Miner{}, err
	// }
	return
}


// func (ser *minerService) FindMiner(tag string, req models.MinerIndexReq) (db *gorm.DB) {
// 	var collect models.Collections
// 	fmt.Printf("tttttt: %s\n", tag)

// 	// db = dao.MinerDao.DB.Where("machine_tag = ?", tag).First(&collect)
// 	db = dao.MinerDao.DB.Table("miner_ironfish").Where("machine_tag = ?", tag).First(&collect)
// 	fmt.Printf("db type: %T", db)
// 	return
// }




