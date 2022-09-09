package category

import (
	"github.com/wpf1118/toolbox/tools/db"
)

type categoryServ struct {
	*db.Mysql
}

func NewCategoryServ() *categoryServ {
	mdb := db.NewMysql()
	return &categoryServ{
		mdb,
	}
}

func (s *categoryServ) AllList() (list []Category, err error) {
	err = s.Order("`order`, `id`").Find(&list).Error
	return
}

func (s *categoryServ) TreeList(list []Category, pid uint) (treeList []TreeCategory) {
	treeList = make([]TreeCategory, 0)
	for _, row := range list {
		if row.Pid == pid {
			t := TreeCategory{
				ID:    row.ID,
				Pid:   row.Pid,
				Name:  row.Name,
				Icon:  row.Icon,
				Order: row.Order,
				Level: row.Level,
			}
			t.Children = s.TreeList(list, row.ID)
			treeList = append(treeList, t)
		}
	}

	return
}
