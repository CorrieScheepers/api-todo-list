package models

type TaskModel struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255);not null"`
	Completed   bool   `gorm:"default:false"`
}
