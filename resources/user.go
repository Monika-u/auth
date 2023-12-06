package resources

type User struct {
	UserId   uint64   `json:"id"`
	Name     string   `json:"name"`
	EmailId  string   `json:"email_id"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

// UserRole represents the possible roles for a user
type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)
