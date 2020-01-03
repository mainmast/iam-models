package user

//User ...
type User struct {
	UUID       string                 `json:"uuid"`
	UserLogin  string                 `json:"user_login"`
	UserSecret string                 `json:"user_secret"`
	IsRoot     bool                   `json:"is_root"`
	UserType   UsrType                `json:"user_type"`
	Status     UsrStatus              `json:"status"`
	CreatedAt  string                 `json:"created_at"`
	UpdatedAt  string                 `json:"updated_at"`
	CustomData map[string]interface{} `json:"custom_data"`
}
