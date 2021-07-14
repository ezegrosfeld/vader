package vader

const (
	//DDD types
	server     = "server"
	repository = "repository"
	service    = "service"
	handler    = "handler"
	domain     = "domain"

	// Generic
	generic = "generic"

	//HTTP Errors
	notFound      = "notFound"
	internal      = "internal"
	badRequest    = "badRequest"
	unauthorized  = "unauthorized"
	alreadyExists = "alreadyExists"

	// HTTP Codes
	notFoundCode      = 404
	internalCode      = 500
	badRequestCode    = 400
	unauthorizedCode  = 401
	alreadyExistsCode = 409
)
