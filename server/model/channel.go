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

// Generate a new passkey for a channel
func (c *Channel) GeneratePassKey() {
	c.PassKey = RandomString(5)
}

// Get One Channel
func (c *Channel) GetChannel(db *gorm.DB, id uint) error {
	err := db.Model(&Channel{}).Preload("Users").Preload("Messages").First(c, id).Error
	return err
}

// Get All Channels
func (c *Channel) GetChannels(db *gorm.DB) (*[]Channel, error) {
	var channels []Channel
	err := db.Model(&Channel{}).Preload("Users").Preload("Messages").Find(&channels).Error
	return &channels, err
}

// Delete one Channel
func (c *Channel) DeleteChannel(db *gorm.DB, id uint) error {
	err := db.Delete(c, id).Error
	return err
}

// Create one Channel
func (c *Channel) CreateChannel(db *gorm.DB) error {
	c.GeneratePassKey()
	err := db.Create(c).Error
	return err
}
