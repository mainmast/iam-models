package resource

// Resource object
type Resource struct {
	UUID        string                 `json:"uuid"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	CreatedAt   string                 `json:"created_at"`
	CustomData  map[string]interface{} `json:"custom_data"`
}
