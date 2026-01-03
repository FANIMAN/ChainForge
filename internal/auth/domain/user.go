package domain

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // hashed
	Role     string `json:"role"`
}
