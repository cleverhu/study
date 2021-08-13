package models

type UserModel struct {
	ID   int    `gorm:"column:user_id;primary_key"`
	Name string `gorm:"column:user_name"`
}

func (this *UserModel) TableName() string {
	return "users"
}
