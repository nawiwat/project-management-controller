package model

type Task struct {
	ID           	uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	ProjectID		uint64			`gorm:"not null" query:"project_id"`
	Name         	string 			`gorm:"not null" json:"name"`
	Description  	string 			`gorm:"not null" json:"description"`
	StartDate    	string 			`gorm:"not null" json:"start_date"`
	EndDate       	string 			`gorm:"not null" json:"end_date"`
	Members			[]string		`gorm:"not null" json:"member"`
	Labels			[]string		`gorm:"not null" json:"labels"`
	Attachments		[]Attachment	`gorm:"not null" json:"attachments"`

	DueDate			string 			`gorm:"not null" json:"due_date"`
	
	Kanban			KanbanColumn		`json:"kanban"`
	Membership   	[]Membership		`gorm:"foreignKey:UserID;references:ID" json:"membership"`
	Notification 	[]Notification		`gorm:"foreignKey:UserID;references:ID" json:"notification"`
}

type KanbanColumn struct {
	ID          uint64 		`gorm:"primarykey,not null;autoIncrement:true;unique" query:"id"`
	ProjectId   uint64		`gorm:"not null"`
	Column 		string		`json:"column"`
	Position 	string		`json:"position"`
}

type Attachment struct {
	ID          uint64 		`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	TaskId		uint64		`gorm:"not null" query:"task_id"`
	Name		string		`gorm:"not null" json:"name"`
	Src			string		`gorm:"not null" json:"src"`
	Size		string		`gorm:"not null" json:"size"`
}

type Comment struct {
	ID          uint64 		`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id" query:"id"`
	TaskId		uint64		`gorm:"not null" query:"task_id"`
	Name		string		`gorm:"not null" json:"name"`
	Src			string		`gorm:"not null" json:"src"`
	Message		string		`gorm:"not null" json:"message"`
	Date		string		`gorm:"not null" json:"date"`
}