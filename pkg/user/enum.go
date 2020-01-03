package user

//UsrType enum,
type UsrType string

const (
	//NoUsrType , no type
	NoUsrType UsrType = ""
	//Plaftorm ..
	Plaftorm UsrType = "platform"
	// API ..
	API UsrType = "api"
)


//UsrStatus enum,
type UsrStatus string

const (
	//NoUsrStatus , no status
	NoUsrStatus UsrStatus = ""
	// Active ..
	Active UsrStatus = "active"
	// Inactive ...
	Inactive UsrStatus = "inactive"
	// Blocked ...
	Blocked UsrStatus = "blocked"
	// Deleted ...
	Deleted UsrStatus = "deleted"
)
