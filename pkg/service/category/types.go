package category

type Category struct {
	ID    uint   `json:"id"`
	Pid   uint   `json:"pid"`
	Name  string `json:"name" gorm:"size:20;comment:分类名称"`
	Icon  string `json:"icon"  gorm:"size:200;comment:分类显示图片"`
	Order int    `json:"order" gorm:"comment: 排序"`
	Level int    `json:"level" gorm:"comment: 等级"`
}

type TreeCategory struct {
	ID       uint           `json:"id"`
	Pid      uint           `json:"pid"`
	Name     string         `json:"name" gorm:"size:20;comment:分类名称"`
	Icon     string         `json:"icon"  gorm:"size:200;comment:分类显示图片"`
	Order    int            `json:"order" gorm:"comment: 排序"`
	Level    int            `json:"level" gorm:"comment: 等级"`
	Children []TreeCategory `json:"children"`
}
