package account

//Account ...
type Account struct {
	UUID       string                 `json:"uuid"`
	Name       string                 `json:"name"`
	IsRoot     bool                   `json:"is_root"`
	CreatedAt  string                 `json:"created_at"`
	UpdatedAt  string                 `json:"updated_at"`
	CustomData map[string]interface{} `json:"custom_data"`
}
