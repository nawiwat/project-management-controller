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
	Membership		[]Membership	
	BoardColumn		[]BoardColumn	
}

type Membership struct {
	ID       		uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	Username		string 			`gorm:"not null" json:"username"`
	Role 			string 			`gorm:"not null" json:"role"`
	JoinDate		time.Time   	`gorm:"type:timestamp;autoCreateTime:true" json:"join_date"`
	UserID			uint64			`json:"user_id"`
	ProjectID		uint64			`json:"project_id"`
	Project			Project			`gorm:"foreignKey:ProjectID;references:ID"`				
}

type Invitation struct {
	ID          	uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	UserID			uint64			`json:"user_id"`
	NotificationID	uint64			`json:"notification_id"`
	ProjectID		uint64			`json:"project_id"`
	Project			Project			`gorm:"foreignKey:ProjectID;references:ID"`
	Status      	string          `json:"status"`
	CreateDate		time.Time   	`gorm:"type:timestamp;autoCreateTime:true" json:"invite_date"`
}