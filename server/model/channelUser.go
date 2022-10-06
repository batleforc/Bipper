package model

import "gorm.io/gorm"

type ChannelUser struct {
	gorm.Model
	ChannelID uint
	UserID    uint
	CanSend   bool
	CanRead   bool
	CanMod    bool
}

// Get all ChannelUsers in channel
func (cu *ChannelUser) GetChannelUsers(db *gorm.DB, id uint) (*[]ChannelUser, error) {
	var channelUsers []ChannelUser
	err := db.Model(&ChannelUser{}).Where("channel_id = ?", id).Find(&channelUsers).Error
	return &channelUsers, err
}

// Get one ChannelUser
func (cu *ChannelUser) GetChannelUser(db *gorm.DB, id uint) error {
	err := db.Model(&ChannelUser{}).First(cu, id).Error
	return err
}

// Delete one ChannelUser
func (cu *ChannelUser) DeleteChannelUser(db *gorm.DB, id uint) error {
	err := db.Delete(cu, id).Error
	return err
}

// Create one ChannelUser
func (cu *ChannelUser) CreateChannelUser(db *gorm.DB) error {
	err := db.Create(cu).Error
	return err
}