package model

type Register struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
}

func (register *Register) TableName() string {
	return "register"
}
