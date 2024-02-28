package model

type Task struct {
	ID           	uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	ProjectId		uint64			`gorm:"not null" json:"project_id" query:"project_id"`
	Name         	string 			`gorm:"not null" json:"name"`
	Description  	string 			`gorm:"not null" json:"description"`
	StartDate    	string 			`gorm:"not null" json:"start_date"`
	EndDate       	string 			`gorm:"not null" json:"end_date"`
	Members			[]TaskMember	`json:"member" gorm:"constraint:OnUpdate:CASCADE;foreignKey:TaskId;references:ID"`
	Attachments		[]Attachment	`json:"attachments" gorm:"constraint:OnUpdate:CASCADE;foreignKey:TaskId;references:ID"`
	Comments		[]Comment		`json:"comments" gorm:"constraint:OnUpdate:CASCADE;foreignKey:TaskId;references:ID"`
	Kanban			KanbanColumn	`json:"kanban"`
}

type TaskMember struct {
	ID				uint64			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	TaskId			uint64 			`json:"task_id"`
	UserId	     	uint64 			`json:"user_id" `
	Username		string			`json:"username" `
}

type KanbanColumn struct {
	ID          	uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	TaskId			uint64 			`json:"task_id"`
	Column 			string			`json:"column" `
	Position 		uint64			`json:"position" `
}

type Attachment struct {
	ID          	uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	TaskId			uint64			`json:"task_id" `
	Name			string			`json:"name" `
	Src				string			`json:"src" `
	Size			string			`json:"size" `
}

type Comment struct {
	ID          	uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	TaskId			uint64			`json:"task_id"`
	Name			string			`json:"name" `
	Src				string			`json:"src" `
	Message			string			`json:"message"`
	Date			string			`json:"date"`
}