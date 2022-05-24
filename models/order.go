package models

type Order struct {
	ID       uint `gorm:"primaryKey" json:"id,omitempty"`
	ItemID   int  `gorm:"not null" json:"item_id" form:"item_id" valid:"required~Your Item id is required for ordering"`
	Quantity int  `gorm:"not null" json:"quantity" form:"description" valid:"required~Your quantity is required ordering"`
	Item     *Item
}
