package image

import (
	"github.com/wpf1118/api/pkg/entity"
	"github.com/wpf1118/toolbox/tools/db"
)

type imageServ struct {
	*db.Mysql
}

func NewImageServ() *imageServ {
	mdb := db.NewMysql()
	return &imageServ{
		mdb,
	}
}

func (s *imageServ) Pics(rid uint) (pics []string, err error) {
	var list []entity.Image
	err = s.Where("rid", rid).Find(&list).Error
	if err != nil {
		return
	}

	for _, v := range list {
		pics = append(pics, v.Url)
	}

	return
}
