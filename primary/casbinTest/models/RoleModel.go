package models

type RoleModel struct {
	ID      int    `gorm:"column:role_id;primary_key"`
	Name    string `gorm:"column:role_name"`
	PID     int    `gorm:"column:role_pid"`
	Comment string `gorm:"column:role_comment"`
}

func (this *RoleModel) TableName() string {
	return "roles"
}
