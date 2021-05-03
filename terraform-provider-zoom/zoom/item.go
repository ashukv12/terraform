package zoom
type Users struct {
	ID            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Type          int    `json:"type"`
	Status        string `json:"status"`
	Pmi           int    `json:"pmi"`
	Time_Zone     string `json:"time_zone"`
	Verified      int    `json:"verified"`
	CreatedAt     string `json:"created_at"`
	LastLoginTime string `json:"last_login_time"`
	PicUrl        string `json:"pic_url"`
	Language      string `json:"language"`
	RoleId        string `json:"role_id"`
	PhoneNumber   string `json:"phone_number"`
}

type UserCreateInfo struct{
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email      string `json:"email"`
	Type       int    `json:"type"`
}
type UserCreate struct {
	Action string `json:"action"`
	UserCreateInfo UserCreateInfo   `json:"user_info"`
}

type Status struct{
	Status string `json:"status"`
}

