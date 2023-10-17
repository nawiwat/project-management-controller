package model

import (
	"time"
)

type Project struct {
	ID       		uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	Name       		string			`json:"name"`
	Description   	string			`gorm:"not null" json:"description"`
	Progess       	string			`json:"progess"`
	Budget      	string			`json:"budget"`
	Deathline      	string			`json:"deathline"`
	Email         	string			`gorm:"not null" json:"email"`
	Github 			string			`gorm:"not null" json:"github"`
	Phone   		string			`gorm:"not null" json:"phone"`
	Membership		[]Membership	

	KanbanBoard		KanbanBoard		`gorm:"foreignKey:ProjectID;references:ID" json:"kanban_board"`
}

type Membership struct {
	Role 		string 			`gorm:"not null" json:"role"`
	JoinDate	time.Time   	`gorm:"type:timestamp;autoCreateTime:true" json:"join_date"`
	UserID		uint64			`json:"user_id"`
	User		User			`gorm:"foreignKey:UserID;references:ID"`
	ProjectID	uint64			`json:"project_id"`
	Project		Project			`gorm:"foreignKey:ProjectID;references:ID"`				
}