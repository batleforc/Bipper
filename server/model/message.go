package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ChannelID uint
	UserID    uint
	Content   string
}
