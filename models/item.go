package models

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"id,omitempty"`
	Name        string `gorm:"not null" json:"name" form:"name" valid:"required~Your Items name is required"`
	Description string `gorm:"not null" json:"description" form:"description" valid:"required~Your Items description is required"`
	Price       int    `gorm:"not null" json:"price" form:"price" valid:"required~Your Items price is required"`
	order       []Order
}
