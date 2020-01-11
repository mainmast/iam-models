package platform

type Status string

const (

	// Platform has no status yet.
	NoStatus Status = ""

	// Platform is pending release - not active.
	Pending Status = "pending"

	// Platform is active
	Active Status = "active"

	// Platform is disabled - no longer active
	Disabled Status = "disabled"

	// Platform is blocked - no longer active
	Blocked Status = "blocked"

	// Platform has been compromised - no longer active
	Compromised Status = "compromised"
)
