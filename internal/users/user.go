package users

type filter struct {
	Id    uint
	Email string
}

type User struct {
	Id       uint
	Name     string
	Email    string
	Password string
}
