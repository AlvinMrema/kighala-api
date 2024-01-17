package utils

type Role int

const (
	Admin Role = iota
	User
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "admin"
	default:
		return "user"
	}
}
