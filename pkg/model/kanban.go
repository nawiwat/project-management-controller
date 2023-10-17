package model

type KanbanBoard struct {
	ID          uint64          `gorm:"primaryKey;autoIncrement:true" json:"id"`
	FraudRuleID uint64          `json:"fraud_rule_id" query:"fraud_rule_id" validate:"required,number"`
	Version     uint64          `json:"version"`
}