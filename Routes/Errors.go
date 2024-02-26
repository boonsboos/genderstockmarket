package routes

type EndpointError struct {
	Kind    string `json:"error"`
	Message string `json:"description"`
}

var ParameterMissing = EndpointError{
	"invalid request",
	"you are missing a required parameter",
}

var LackingResources = EndpointError{
	"invalid request",
	"you do not have the resources to complete this action",
}
