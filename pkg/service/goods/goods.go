package goods

import (
	"github.com/wpf1118/api/pkg/dto"
	"github.com/wpf1118/api/pkg/entity"
	"github.com/wpf1118/toolbox/tools/db"
	"github.com/wpf1118/toolbox/tools/logging"
	"gorm.io/gorm"
)

type goodsServ struct {
	*db.Mysql
}

func NewGoodsServ() *goodsServ {
	mdb := db.NewMysql()
	return &goodsServ{
		mdb,
	}
}

func (s *goodsServ) Paginate(req Req) (count int64, list []entity.Goods, err error) {
	tx := s.Model(&entity.Goods{}).
		Scopes(
			dto.MakeCondition(req),
			dto.Paginate(req.GetSize(), req.GetPage()),
		).
		Find(&list).Count(&count)

	//sql := s.ToSQL(func(tx *gorm.DB) *gorm.DB {
	//	return tx.Model(&Goods{}).
	//		Scopes(
	//			dto.MakeCondition(req),
	//			dto.Paginate(req.GetSize(), req.GetPage()),
	//		).
	//		Find(&list).Count(&count)
	//})
	//
	//logging.DebugF("%s", sql)

	err = tx.Error
	if err != nil {
		logging.ErrorF("db error: %s", err)
		return
	}

	return
}

func (s *goodsServ) Add(req ReqAdd) (id uint, err error) {
	goods := &entity.Goods{
		Name:          req.Name,
		Cid:           req.Cid,
		Pic:           req.Pic,
		Status:        req.Status,
		OriginalPrice: req.OriginalPrice,
		SellPrice:     req.SellPrice,
		Stock:         req.Stock,
	}

	err = s.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(goods).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		for _, pic := range req.GoodsPics {
			if err := tx.Create(&entity.Image{
				Rid: goods.ID,
				Url: pic,
			}).Error; err != nil {
				return err
			}
		}

		// 返回 nil 提交事务
		return nil
	})

	if err != nil {
		return
	}

	id = goods.ID

	return
}

func (s *goodsServ) Detail(id uint) (goods entity.Goods, err error) {
	err = s.First(&goods, id).Error
	return
}
