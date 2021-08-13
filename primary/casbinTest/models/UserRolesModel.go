package models

type UserRolesModel struct {
	UserID int    `gorm:"column:user_id"`
	RoleID string `gorm:"column:role_id"`
}

func (this *UserRolesModel) TableName() string {
	return "user_roles"
}
