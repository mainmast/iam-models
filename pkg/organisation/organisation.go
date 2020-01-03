package organisation

import "github.com/mainmast/iam-models/pkg/user"

//Organisation ...
type Organisation struct {
	UUID             string                 `json:"uuid"`
	Name             string                 `json:"name"`
	Status           OrgStatus              `json:"status"`
	CreatedAt        string                 `json:"created_at"`
	UpdatedAt        string                 `json:"updated_at"`
	WhitelistDomains []string               `json:"whitelist_domains"`
	CustomData       map[string]interface{} `json:"custom_data"`
}

// CreateOrgRQ ...
type CreateOrgRQ struct {
	Organisation Organisation `json:"organisation"`
	User         user.User    `json:"user"`
}

//CreateOrgRS ...
type CreateOrgRS struct {
	OrganisationName  string `json:"organisation_name"`
	AccountName       string `json:"account_name"`
	UserPlatformLogin string `json:"user_platform_login"`
	UserAPILogin      string `json:"user_api_login"`
	OrganisationUUID  string `json:"organisation_uuid"`
	UserPlatformUUID  string `json:"user_platform_uuid"`
	UserAPIUUID       string `json:"user_api_uuid"`
	AccountUUID       string `json:"account_uuid"`
	AccessUUID        string `json:"access_uuid"`
}
