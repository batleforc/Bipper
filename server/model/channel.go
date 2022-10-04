package model

import "gorm.io/gorm"

type Channel struct {
	gorm.Model
	Name        string
	Description string
	Picture     string
	Private     bool
	PassKey     string // crypted field (can be regenerated) and empty if public
	Users       []ChannelUser
	Owner       uint
	Messages    []Message
}

type ChannelUser struct {
	gorm.Model
	ChannelID uint
	UserID    uint
	CanSend   bool
	CanRead   bool
	CanMod    bool
}
