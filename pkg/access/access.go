package access

// Access key
type Key struct {
	UUID      string `json:"uuid"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	CreatedAt string `json:"created_at"`
	LastUsed  string `json:"last_used"`
	Status    Status `json:"status"`
	UserUUID  int64  `json:"user_uuid"`
	ExpiresAt string `json:"expires_at"`
}
