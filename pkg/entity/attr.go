package entity

type Attr struct {
	Base
	Key    string `json:"key" gorm:"size:50;comment:属性key"`
	Name   string `json:"name" gorm:"size:50;comment:属性名称"`
	Values string `json:"values" gorm:"size:500;comment:属性可选值，多个值之间用','分隔"`
}
