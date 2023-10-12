package responses

type ResponseUsers struct {
	ID            	uint64 			`json:"id"`
	Username      	string			`json:"username"`
	Password       	string			`json:"password"`
	Name  			string			`json:"name"`
	Surname  		string			`json:"surname"`
	Email         	string			`json:"email"`
	Github 			string			`json:"github"`
	Phone   		string			`json:"phone"`
	Description   	string			`json:"description"`
}