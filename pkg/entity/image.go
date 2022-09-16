package entity

type Image struct {
	Base
	Url   string `json:"url" gorm:"size:100;comment:图片地址"`
	Rid   uint   `json:"rid" gorm:"comment:关联id, sku_id或spu_id或其它"`
	Order int    `json:"order" gorm:"comment:排序"`
}
