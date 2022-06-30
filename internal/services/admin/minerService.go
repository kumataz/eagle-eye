/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-14 13:29:28
 */
package admin

import (
	// "strings"

	"github.com/kumataahh/eagle-eye/internal/dao"
	"github.com/kumataahh/eagle-eye/internal/models"

	"gorm.io/gorm"
)

type minerService struct{}

var MinerService = minerService{}

func (ser *minerService) GetMiner(minerId uint) (miner models.Miner, err error) {
	miner.MinerId = minerId
	err = dao.MinerDao.DB.First(&miner).Error
	if err != nil {
		return models.Miner{}, err
	}
	return
}

func (ser *minerService) GetMiners(req models.MinerIndexReq) (db *gorm.DB) {

	db = dao.AuDao.DB.Table("miner_ironfish")

	// if req.Title != "" {
	// 	db = db.Where("title like ?", "%"+req.Title+"%")
	// }

	// if req.CreatedAt != "" {
	// 	period := strings.Split(req.CreatedAt, " ~ ")
	// 	start := period[0] + " 00:00:00"
	// 	end := period[1] + " 23:59:59"
	// 	db = db.Where("created_at > ? ", start).Where("created_at < ?", end)
	// }

	return
}

// //添加或保存文章信息
// func (ser *minerService) SaveArticle(req models.ArticleReq) (err error) {

// 	var (
// 		article models.Article
// 	)

// 	if req.ArticleId > 0 {

// 		article.ArticleId = uint(req.ArticleId)

// 		dao.ArticleDao.DB.First(&article)

// 		article.Title = req.Title
// 		article.Desc = req.Desc
// 		article.Content = req.Content
// 		article.CoverImg = req.CoverImg

// 		err = dao.ArticleDao.DB.Save(&article).Error
// 		if err != nil {
// 			return
// 		}

// 	} else {
// 		article.Title = req.Title
// 		article.Content = req.Content
// 		article.CoverImg = req.CoverImg
// 		article.Desc = req.Desc
// 		err = dao.ArticleDao.DB.Save(&article).Error
// 		if err != nil {
// 			return
// 		}
// 	}

// 	return
// }

// func (ser *minerService) DelArticle(id int) (err error) {
// 	var article models.Article
// 	article.ArticleId = uint(id)
// 	err = dao.ArticleDao.DB.Delete(&article).Error
// 	return
// }
