package model

type Customer struct {
	ID      uint   `gorm:"primarykey" json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
