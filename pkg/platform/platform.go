package platform

// Platform Object
type Platform struct {
	UUID             string                 `json:"uuid"`
	Name             string                 `json:"name"`
	Status           Status                 `json:"status"`
	CreatedAt        string                 `json:"created_at"`
	UpdatedAt        string                 `json:"updated_at"`
	CustomData       map[string]interface{} `json:"custom_data"`
}
