package organisation

// Organisation Status Enum
type Status string

const (
	// No status
	NoStatus Status = ""

	// Organisation is pending activation - not yet active.
	Pending Status = "pending"

	// Organisation is active
	Active Status = "active"

	// The organisation has been disabled - not active.
	Disable Status = "disable"

	// The organisation has been blocked - not active.
	Blocked Status = "blocked"

	// The organisation has been compromised - not active.
	Compromised Status = "compromised"
)
