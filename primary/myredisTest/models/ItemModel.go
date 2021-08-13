package models

type ItemModel struct {
	ID   int    `gorm:"column:id;primary_key" json:"id"`
	Type string `gorm:"column:type;" json:"type"`
}

func NewItemModel() *ItemModel {
	return &ItemModel{}
}

func (this *ItemModel) TableName() string {
	return "items"
}
