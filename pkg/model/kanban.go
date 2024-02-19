package model

type BoardColumn struct {
	ID          	uint64          `gorm:"primaryKey;autoIncrement:true" json:"id"`
	ProjectID		uint64			`json:"ProjectID"`
	Postiton 		uint64			`json:"position"`
	Name			string			`json:"name"`
}

