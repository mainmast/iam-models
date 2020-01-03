package permission

//Permission ....
type Permission struct {
	UUID        string                 `json:"uuid"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Resource    int                    `json:"resource"`
	CreatedAt   string                 `json:"created_at"`
	CustomData  map[string]interface{} `json:"custom_data"`
}
