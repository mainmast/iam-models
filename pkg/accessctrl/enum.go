package accessctrl

//TokenStatus enum,
type TokenStatus string

const (
	//NoTknStatus , no status
	NoTknStatus TokenStatus = ""
	// Active ..
	Active TokenStatus = "active"
	// Expired ...
	Expired TokenStatus = "expired"
	// Revoked ...
	Revoked TokenStatus = "revoked"
	// Renewed ...
	Renewed TokenStatus = "renewed"
)
