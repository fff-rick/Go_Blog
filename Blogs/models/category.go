package models

type Category struct {
	Cid      int
	Name     string
	CreateAt string
	UpdateAt string
}

// TableName 指定表名
func (Category) TableName() string {
	return "category"
}

type CategoryResponse struct {
	*HomeResponse
	CategoryName string
}
