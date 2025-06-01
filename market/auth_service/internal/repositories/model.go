package repo

type User struct {
	ID           uint
	Login        string
	PasswordHash string
	Role         string
}
