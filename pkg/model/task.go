package model

type Task struct {
	ID           	uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	ProjectId		uint64			`gorm:"not null" query:"project_id"`
	Name         	string 			`gorm:"not null" json:"name"`
	Description  	string 			`gorm:"not null" json:"description"`
	StartDate    	string 			`gorm:"not null" json:"start_date"`
	EndDate       	string 			`gorm:"not null" json:"end_date"`
	Members			[]TaskMember	`json:"member"`
	Attachments		[]Attachment	`json:"attachments"`
	Comments		[]Comment		`json:"comments"`
	DueDate			string 			`gorm:"not null" json:"due_date"`
	Kanban			KanbanColumn	`json:"kanban"`
}

type TaskMember struct {
	ID				uint64			`gorm:"primarykey,not null;autoIncrement:true;unique" query:"id"`
	TaskId			uint64 			`json:"task_id"`
	UserId	     	uint64 			`json:"user_id"`
	Username		string			`json:"username"`
}

type KanbanColumn struct {
	ID          	uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" query:"id"`
	TaskId			uint64 			`json:"task_id"`
	Column 			string			`json:"column"`
	Position 		uint64			`json:"position"`
}

type Attachment struct {
	ID          	uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	TaskId			uint64			`gorm:"not null" query:"task_id"`
	Name			string			`gorm:"not null" json:"name"`
	Src				string			`gorm:"not null" json:"src"`
	Size			string			`gorm:"not null" json:"size"`
}

type Comment struct {
	ID          	uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	TaskId			uint64			`gorm:"not null" query:"task_id"`
	Name			string			`gorm:"not null" json:"name"`
	Src				string			`gorm:"not null" json:"src"`
	Message			string			`gorm:"not null" json:"message"`
	Date			string			`gorm:"not null" json:"date"`
}