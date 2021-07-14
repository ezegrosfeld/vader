package vader

// Returns a generic error
func Generic(msg string) Error {
	e := err{msg: msg, code: 0, group: generic, kind: generic}
	return &e
}

// Returns a not found error
func NotFound(msg string) Error {
	e := err{msg: msg, code: notFoundCode, group: generic, kind: notFound}
	return &e
}

// Returns a bad request error
func BadRequest(msg string) Error {
	e := err{msg: msg, code: badRequestCode, group: generic, kind: badRequest}
	return &e
}

// Returns an internal error
func InternalError(msg string) Error {
	e := err{msg: msg, code: internalCode, group: generic, kind: internal}
	return &e
}

// Returns an alredy exists error
func AlredyExistsError(msg string) Error {
	e := err{msg: msg, code: alreadyExistsCode, group: generic, kind: alreadyExists}
	return &e
}

// Returns an unauthorized error
func Unauthorized(msg string) Error {
	e := err{msg: msg, code: unauthorizedCode, group: generic, kind: unauthorized}
	return &e
}

// Returns a new error with status code
func NewError(msg string, status int) Error {
	e := err{msg: msg, code: status, group: generic, kind: generic}
	return &e
}
