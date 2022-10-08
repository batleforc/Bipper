package model

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model
	Name        string
	Description string
	Picture     string
	Private     bool
	PassKey     string `json:"-"` // crypted field (can be regenerated) and empty if public
	Users       []ChannelUser
	Owner       uint
	Messages    []Message
}

// Generate a new passkey for a channel
func (c *Channel) GeneratePassKey() string {
	pass := RandomString(5)
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	c.PassKey = string(hash)
	return pass
}

// Check if the passkey is correct
func (c *Channel) CheckPassKey(pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(c.PassKey), []byte(pass))
	return err == nil
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

// Get All public Channels
func (c *Channel) GetPublicChannels(db *gorm.DB, limit, page int) (*[]Channel, error) {
	var channels []Channel
	err := db.Model(&Channel{}).Preload("Users").Preload("Messages").Offset((page-1)*limit).Limit(limit).Where("private = ?", false).Find(&channels).Error
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

// Update one Channel
func (c *Channel) UpdateChannel(db *gorm.DB) error {
	err := db.Save(c).Error
	return err
}

// Get paginated channel
func (c *Channel) GetPaginatedChannel(db *gorm.DB, limit, page int) []Channel {
	var channels []Channel
	db.Model(&Channel{}).Preload("Messages").Preload("Users").Offset((page - 1) * limit).Limit(limit).Find(&channels)
	return channels
}

// Get paginated public channel
func (c *Channel) GetPaginatedPublicChannel(db *gorm.DB, limit, page int) []Channel {
	var channels []Channel
	db.Model(&Channel{}).Preload("Messages").Preload("Users").Where("private = ?", false).Offset((page - 1) * limit).Limit(limit).Find(&channels)
	return channels
}
