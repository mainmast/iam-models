package organisation

//OrgStatus enum,
type OrgStatus string

const (
	//NoOrgStatus , no status
	NoOrgStatus OrgStatus = ""
	//Pending ..
	Pending OrgStatus = "pending"
	// Active ..
	Active OrgStatus = "active"
	// Disable ...
	Disable OrgStatus = "disable"
	// Blocked ...
	Blocked OrgStatus = "blocked"
)
