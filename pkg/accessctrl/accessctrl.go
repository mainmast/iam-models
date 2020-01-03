package accessctrl

// AccessCtrl ...
type AccessCtrl struct {
	UUID      string      `json:"uuid"`
	AccessKey string      `json:"access_key"`
	SecretKey string      `json:"secret_key"`
	CreatedAt string      `json:"created_at"`
	LastUsed  string      `json:"last_used"`
	Status    TokenStatus `json:"status"`
	UserUUID  int64       `json:"user_uuid"`
	ExpiresAt string      `json:"expires_at"`
}
