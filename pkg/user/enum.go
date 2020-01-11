package user

// User Status Enum
type Status string

const (
	// No status
	NoStatus Status = ""

	// User is pending activation - not active.
	Pending Status = "pending"

	// User is active
	Active Status = "active"

	// User is deemed inactive - no longer active.
	Inactive Status = "inactive"

	// User has been disabled - no longer active.
	Disabled Status = "disabled"

	// User has been compromised - no longer active.
	Compromised Status = "compromised"

	// User has been deleted - no longer active.
	Deleted Status = "deleted"
)
