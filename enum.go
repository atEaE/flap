package valigo

type requiredType string

var (
	RequiredAllowEmpty               = requiredType("allowEmpty")
	RequiredDeniedEmpty              = requiredType("deniedEmpty")
	RequiredDeniedEmptyWithTrimspace = requiredType("deniedEmptyWithTrimSpace")
)
