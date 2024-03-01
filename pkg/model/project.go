package model

import (
	"time"
)

type Project struct {
	ID       		uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	Name       		string			`json:"name"`
	Description   	string			`json:"description"`
	Progess       	string			`json:"progess"`
	Budget      	string			`json:"budget"`
	Deathline      	string			`json:"deathline"`
	Email         	string			`json:"email"`
	Github 			string			`json:"github"`
	Phone   		string			`json:"phone"`
	Invitation		[]Invitation	`json:"invitation"`
	Task			[]Task			`json:"task"`
	Membership		[]Membership	`json:"membership"`
}

type Membership struct {
	ID       		uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	Username		string 			`gorm:"not null" json:"username"`
	Role 			string 			`gorm:"not null" json:"role"`
	JoinDate		time.Time   	`gorm:"type:timestamp;autoCreateTime:true" json:"join_date"`
	UserId			uint64			`json:"user_id"`
	ProjectId		uint64			`json:"project_id"`				
}

type Invitation struct {
	ID          	uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	UserId			uint64			`json:"user_id"`
	ProjectId		uint64			`json:"project_id"`
	Status      	string          `json:"status"`
	CreateDate		time.Time   	`gorm:"type:timestamp;autoCreateTime:true" json:"invite_date"`
}

type InviteReq struct {
	UserId			uint64		`json:"user_id"`
	ProjectId		uint64		`json:"project_id"`
	Sender 			string		
}