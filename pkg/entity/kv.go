package entity

type Kv struct {
	Base
	Key   string `json:"key"`
	Value string `json:"value"`
}
