package auth

// AuthRole ...
type AuthRole string

const (
	Admin AuthRole = "seller"
	User  AuthRole = "buyer"
)

//String convert role enum to string
func (a AuthRole) String() string {
	switch a {
	case Admin:
		return "admin"
	case User:
		return "user"
	default:
		return "invalid"
	}
}
