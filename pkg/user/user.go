package user

// User Object
type User struct {
	UUID           string                 `json:"uuid"`
	UserLogin      string                 `json:"user_login"`
	UserSecret     string                 `json:"user_secret"`
	IsRoot         bool                   `json:"is_root"`
	PlatformAccess bool                   `json:"platform_access"`
	APIAccess      bool                   `json:"api_access"`
	Status         Status                 `json:"status"`
	CustomData     map[string]interface{} `json:"custom_data"`
	CreatedAt      string                 `json:"created_at"`
	UpdatedAt      string                 `json:"updated_at"`
}
