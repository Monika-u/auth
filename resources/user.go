package resources

type User struct {
	UserId   uint64 `json:"user_id"`
	Name     string `json:"name"`
	EmailId  string `json:"email_id"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}
