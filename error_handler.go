package vader

// NewRepoErrorHandler returns an error handler with group of repo
func NewRepoErrorHandler() *errorHandler {
	group := repository
	return &errorHandler{group}
}

// NewServiceErrorHandler returns an error handler with group of service
func NewServiceErrorHandler() *errorHandler {
	group := service
	return &errorHandler{group}
}

// NewDomainErrorHandler returns an error handler with group of domain
func NewDomainErrorHandler() *errorHandler {
	group := domain
	return &errorHandler{group}
}

// NewHandlerErrorHandler returns an error handler with group of handler
func NewHandlerErrorHandler() *errorHandler {
	group := handler
	return &errorHandler{group}
}

func (eh *errorHandler) Generic(msg string) Error {
	e := new(err)
	e.code = 0
	e.msg = msg
	e.group = eh.group
	e.kind = generic
	return e
}

func (eh *errorHandler) NotFound(msg string) Error {
	e := new(err)
	e.code = notFoundCode
	e.msg = msg
	e.group = eh.group
	e.kind = notFound
	return e
}

func (eh *errorHandler) InternalError(msg string) Error {
	e := new(err)
	e.code = internalCode
	e.msg = msg
	e.group = eh.group
	e.kind = internal
	return e
}

func (eh *errorHandler) BadRequest(msg string) Error {
	e := new(err)
	e.code = badRequestCode
	e.msg = msg
	e.group = eh.group
	e.kind = badRequest
	return e
}

func (eh *errorHandler) Unauthorized(msg string) Error {
	e := new(err)
	e.code = unauthorizedCode
	e.msg = msg
	e.group = eh.group
	e.kind = unauthorized
	return e
}

func (eh *errorHandler) AlreadyExists(msg string) Error {
	e := new(err)
	e.code = alreadyExistsCode
	e.msg = msg
	e.group = eh.group
	e.kind = alreadyExists
	return e
}

func (eh *errorHandler) UnprocessableEntity(msg string) Error {
	e := new(err)
	e.code = unprocessableEntityCode
	e.msg = msg
	e.group = eh.group
	e.kind = unprocessableEntity
	return e
}
