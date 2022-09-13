package service

type Pagination struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

func (m *Pagination) GetPage() int {
	if m.Page <= 0 {
		m.Page = 1
	}
	return m.Page
}

func (m *Pagination) GetSize() int {
	if m.Size <= 0 {
		m.Size = 10
	}
	return m.Size
}
