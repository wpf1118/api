package goods

import (
	"github.com/wpf1118/api/pkg/dto"
	"github.com/wpf1118/toolbox/tools/db"
	"github.com/wpf1118/toolbox/tools/logging"
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

func (s *goodsServ) Paginate(req Req) (count int64, list []Goods, err error) {
	tx := s.Model(&Goods{}).
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
