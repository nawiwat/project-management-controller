package model

type User struct {
	ID          	uint64 			`gorm:"primarykey,not null;autoIncrement:true;unique" json:"id"`
	Username      	string			`gorm:"not null" json:"username"`
	Password       	string			`gorm:"not null" json:"password"`
	Name  			string			`gorm:"not null" json:"name"`
	Surname  		string			`gorm:"not null" json:"surname"`
	Email         	string			`gorm:"not null" json:"email"`
	Github 			string			`gorm:"not null" json:"github"`
	Phone   		string			`gorm:"not null" json:"phone"`
	Description   	string			`gorm:"not null" json:"description"`
	Membership		[]Membership
}

