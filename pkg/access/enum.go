package access

// Token status Enum
type Status string

const (
	// The token does not have a status
	NoTknStatus Status = ""

	// Token is currently active and valid
	Active Status = "active"

	// Token has expired and is no longer valid
	Expired Status = "expired"

	// Token has been revoked and is no longer valid.
	Revoked Status = "revoked"

	// Token has been renewed and is valid.
	Renewed Status = "renewed"
)
