package model

type KanbanBoard struct {
	ID          uint64          `gorm:"primaryKey;autoIncrement:true" json:"id" `
	ProjectID 	uint64          `json:"project_id" query:"project_id" validate:"number"`
	Column		[]BoardColumn	`gorm:"foreignKey:KanbanBoardID;references:ID" json:"columns"`
}

type BoardColumn struct {
	ID          uint64          `gorm:"primaryKey;autoIncrement:true" json:"id"`
	KanbanBoardID	uint64		`json:"kanban_board_id" query:"kanban_board_id" validate:"number"`
	Name		string			`json:"name"`
}