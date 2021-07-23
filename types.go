package vader

// Error interface
type Error interface {
	// Error returns the string representation of the error
	Error() string
	// Checks if the error has group of server
	IsServer() bool
	// Checks if the error has group of repository
	IsRepository() bool
	// Checks if the error has group of service
	IsService() bool
	// Checks if the error has group of handler
	IsHandler() bool
	// Checks if the error has group of domain
	IsDomain() bool
	// Checks if the error has kind of not found
	IsNotFound() bool
	// Checks if the error has kind of internal
	IsInternal() bool
	// Checks if the error has kind of bad request
	IsBadRequest() bool
	// Checks if the error has kind of alredy exists
	IsAlredyExists() bool
	// Returns the group of the error
	GetGroup() string
	// Returns the kind of the error
	GetKind() string
	// Returns the code of the error
	GetCode() int
}

// Error is the base error
type err struct {
	// Message of the error
	msg string
	// Group of the error
	group string
	// Kind of the error
	kind string
	// Code of the error
	code int
}

// ErrorHandler has a pre defined group
type ErrorHadler interface {
	// Returns a generic error with the group of the handler
	Generic(msg string) Error
	// Returns an error with the group of the handler and kind of not found
	NotFound(msg string) Error
	// Returns an error with the group of the handler and kind of internal
	Internal(msg string) Error
	// Returns an error with the group of the handler and kind of bad request
	BadRequest(msg string) Error
	// Returns an error with the group of the handler and kind of unauthorized
	Unauthorized(msg string) Error
	// Returns an error with the group of the handler and kind of alredy exists
	AlreadyExists(msg string) Error
	// Returns an error with the group of the handler and kind of unprocessable entity
	UnprocessableEntity(msg string) Error
}

// Struct of handler containing gruop
type errorHandler struct {
	// Group of the errors
	group string
}
