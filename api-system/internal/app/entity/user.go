package entity

type UserRegister struct {
	UserName    string
	Email       string
	Phone       string
	Password    string
	InvitedCode string
}

type UserLogin struct {
	UserName string
	Password string
}
