package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ChannelID uint
	UserID    uint
	Content   string
}

// Get all Messages in channel
func (m *Message) GetMessages(db *gorm.DB, id uint) (*[]Message, error) {
	var messages []Message
	err := db.Model(&Message{}).Where("channel_id = ?", id).Find(&messages).Error
	return &messages, err
}

// Create one Message
func (m *Message) CreateMessage(db *gorm.DB) error {
	err := db.Create(m).Error
	return err
}

// Delete one Message
func (m *Message) DeleteMessage(db *gorm.DB, id uint) error {
	err := db.Delete(m, id).Error
	return err
}

// Update one Message
func (m *Message) UpdateMessage(db *gorm.DB, id uint) error {
	err := db.Model(m).Where("id = ?", id).Updates(m).Error
	return err
}
