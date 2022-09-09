package entity

const (
	CategoryLevelOne  = 1
	CategoryLevelTwo  = 2
	CategoryLevelTree = 3
)

type Category struct {
	Base
	Pid   uint   `json:"pid"`
	Name  string `json:"name" gorm:"size:20;comment:分类名称"`
	Icon  string `json:"icon"  gorm:"size:200;comment:分类显示图片"`
	Order int    `json:"order" gorm:"comment: 排序"`
	Level int    `json:"level" gorm:"comment: 等级"`
}
