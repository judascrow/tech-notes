package models

type Note struct {
	ID        uint   `json:"id" gorm:"primaryKey;type:int(11)"`
	Title     string `json:"title" gorm:"type:varchar(50);not null"`
	Text      string `json:"text" gorm:"type:varchar(500);not null"`
	Completed bool   `json:"completed" gorm:"default:false"`
	BaseModel
}
