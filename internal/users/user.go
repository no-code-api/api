package users

type User struct {
	Id       uint
	Name     string
	Email    string
	Password string
}

type UserFilter struct {
	Id    uint
	Email string
}
