package entity

type Goods struct {
	Base
	Name string `json:"name" gorm:"size:20;comment:商品名称"`
}
