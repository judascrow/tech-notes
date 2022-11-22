package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;type:int(11)"`
	Username string `json:"username" gorm:"type:varchar(50);unique_index;not null"`
	Password string `json:"-" gorm:"type:varchar(100);not null"`
	Active   bool   `json:"active" gorm:"default:true"`
	RoleID   uint   `json:"-" gorm:"type:int(11)"`
	Role     Role   `json:"role"`

	BaseModel
}

type Role struct {
	ID   uint   `json:"id" gorm:"primaryKey;type:int(11)"`
	Name string `json:"name" gorm:"type:varchar(50);not null"`
}
