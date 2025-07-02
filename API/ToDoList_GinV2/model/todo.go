package model

type Todo struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Task string `json:"task"`
}
