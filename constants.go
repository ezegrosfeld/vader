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
	notFound            = "notFound"
	internal            = "internal"
	badRequest          = "badRequest"
	unauthorized        = "unauthorized"
	alreadyExists       = "alreadyExists"
	unprocessableEntity = "unprocessableEntity"
	forbidden           = "forbidden"
	conflict            = "conflict"

	// HTTP Codes
	notFoundCode            = 404
	internalCode            = 500
	badRequestCode          = 400
	unauthorizedCode        = 401
	alreadyExistsCode       = 409
	unprocessableEntityCode = 422
	forbiddenCode           = 403
	conflictCode            = 409
)
