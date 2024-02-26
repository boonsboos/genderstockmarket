package auth

type AuthError struct {
	Kind    string `json:"error"`
	Message string `json:"description"`
}

func (e AuthError) Error() string {
	return e.Message
}

var (
	MalformattedAuthorizationHeader = AuthError{
		"unauthorized",
		"malformatted authorization header",
	}
	InternalServerError = AuthError{
		"internal server error",
		"this is not your fault, try again later.",
	}
	Unauthorized = AuthError{
		"unauthorized",
		"you do not have permission to view this page",
	}
)
