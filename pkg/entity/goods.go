package entity

type Goods struct {
	Base
	Name string `json:"name" gorm:"size:20;comment:商品名称"`
	Cid  uint   `json:"cid" gorm:"comment:商品分类"`
	Pic  string `json:"pic" gorm:"size:100;comment:商品主图"`
}
