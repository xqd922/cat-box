package models

type Options struct {
	ID          int8   `gorm:"primaryKey" json:"-"`
	Name        string `gorm:"unique;default:options" json:"-"`
	UpdateDelay string `gorm:"default:30m" json:"update_delay" form:"update_delay"`
}
