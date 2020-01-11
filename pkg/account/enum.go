package account

// Account status Enum
type Status string

const (
	// The account does not have a status
	NoStatus Status = ""

	// The account is currently active.
	Active Status = "active"

	// The account is pending activation - not yet active.
	Pending Status = "pending"

	// The account activation has expired - a new activation process will need to be followed.
	Expired Status = "expired"

	// The account has been disabled.
	Disabled Status = "disabled"

	// The account has been compromised - this is disabled, and separated for additional triggers.
	Compromised Status = "compromised"
)
