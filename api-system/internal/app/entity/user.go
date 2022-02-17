package entity

type UserRegister struct {
	UserName string
	Email    string
	Phone    string
	Password string
}

type UserLogin struct {
	UserName string
	Password string
}
