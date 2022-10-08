package model

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string

const (
	Admin  Role = "Admin"
	Modo   Role = "Modo"
	Member Role = "Member"
)

type PublicUser struct {
	Email   string
	Pseudo  string
	Picture string
	Name    string
	Surname string
	Role    Role
}

type User struct {
	gorm.Model
	Email      string
	Pseudo     string
	Picture    string
	Name       string
	Surname    string
	Password   string `json:"-"`
	Role       Role
	Tokens     []Token   `json:"-"`
	MyChannels []Channel `gorm:"foreignKey:Owner;"`
	Channels   []ChannelUser
	// Add push notification body
}

func (u *User) ToPublicUser() PublicUser {
	return PublicUser{
		Email:   u.Email,
		Pseudo:  u.Pseudo,
		Picture: u.Picture,
		Name:    u.Name,
		Surname: u.Surname,
		Role:    u.Role,
	}
}

func (u *User) HashPassword(pass string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	u.Password = string(hash)
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
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

// Get One User By Mail
func (u *User) GetUserByMail(db *gorm.DB, mail string) error {
	err := db.Model(&User{}).Preload("Tokens").Preload("Channels").Preload("MyChannels").Where("email = ?", mail).First(u).Error
	return err
}

// Get One User By Pseudo
func (u *User) GetUserByPseudo(db *gorm.DB, pseudo string) error {
	err := db.Model(&User{}).Preload("Tokens").Preload("Channels").Preload("MyChannels").Where("pseudo = ?", pseudo).First(u).Error
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
func (u *User) UpdateOrCreateUser(db *gorm.DB) error {
	err := db.Save(u).Error
	return err
}
