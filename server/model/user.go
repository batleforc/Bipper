package model

import (
	"gorm.io/gorm"
)

type Role string

const (
	Admin  Role = "Admin"
	Modo   Role = "Modo"
	Member Role = "Member"
)

type User struct {
	gorm.Model
	Email      string
	Pseudo     string
	Picture    string
	Name       string
	Surname    string
	Password   string
	Role       Role
	Tokens     []Token
	MyChannels []Channel `gorm:"foreignKey:Owner;"`
	Channels   []ChannelUser
	// Add push notification body
}

func (u *User) IsAdmin() bool {
	return u.Role == Admin
}

func (u *User) IsModo() bool {
	return u.Role == Modo
}

func (u *User) IsMember() bool {
	return u.Role == Member
}

// Get One User
func (u *User) GetUser(db *gorm.DB, id uint) error {
	err := db.Model(&User{}).Preload("Tokens").Preload("Channels").Preload("MyChannels").First(u, id).Error
	return err
}

// Get All Users
func (u *User) GetUsers(db *gorm.DB) (*[]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("Tokens").Preload("Channels").Preload("MyChannels").Find(&users).Error
	return &users, err
}

// Delete one User
func (u *User) DeleteUser(db *gorm.DB, id uint) error {
	err := db.Delete(u, id).Error
	return err
}

// Update or create one user
func (u *User) UpdateOrCreateUser(db *gorm.DB, id uint) error {
	err := db.Save(u).Error
	return err
}
