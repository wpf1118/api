package entity

type Goods struct {
	Base
	Name          string `json:"name" gorm:"size:20;comment:商品名称"`
	Cid           uint   `json:"cid" gorm:"comment:商品分类"`
	Pic           string `json:"pic" gorm:"size:100;comment:商品主图"`
	Status        int    `json:"status" gorm:"default:0;comment:0 未上架 1 正常 2 下架"`
	OriginalPrice uint   `json:"original_price" gorm:"default:9999999;comment:原价 单位分"`
	SellPrice     uint   `json:"sell_price" gorm:"default:9999999;comment:售卖价 单位分"`
	Stock         uint   `json:"stock" gorm:"default:0;comment:库存"`
}
