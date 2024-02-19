package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID           uint64 `gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	Username     string `gorm:"not null" json:"username"`
	Password     string `gorm:"not null" json:"password"`
	Name         string `gorm:"not null" json:"name"`
	Surname      string `gorm:"not null" json:"surname"`
	Email        string `gorm:"not null" json:"email"`
	Github       string `gorm:"not null" json:"github"`
	Phone        string `gorm:"not null" json:"phone"`
	Description  string `gorm:"not null" json:"description"`
	Attachment   ProfileAttachment      `gorm:"constraint:OnUpdate:CASCADE" json:"attachment"`
	Membership   []Membership			`json:"membership"`
	Notification []Notification			`json:"notification"`
}

type Notification struct {
	ID          	uint64 `gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	UserID      	uint64 `json:"user_id"`
	SendBy			string `gorm:"not null" json:"send_by"`
	Type        	string `gorm:"not null" json:"type"`
	Description 	string `gorm:"not null" json:"description"`
	InviteId    	uint64 `json:"invite_id"`
	TaskId			uint64 `json:"task_id"`
	Date			time.Time   	`gorm:"type:timestamp;autoCreateTime:true" json:"date"`
}

type UserLogin struct {
	Username 	string `gorm:"not null" json:"username"`
	Password 	string `gorm:"not null" json:"password"`
}

type UserToken struct {
	ID       	uint64 `gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	Username 	string `json:"username"`
	jwt.StandardClaims
}

type ProfileAttachment struct {
	ID       	uint64 		`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	UserId		uint64		`gorm:"not null" json:"user_id"`
	Name		string		`gorm:"not null" json:"name"`
	Src			string		`gorm:"not null" json:"src"`
	Size		string		`gorm:"not null" json:"size"`
}