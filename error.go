package vader

func (e *err) Error() string {
	return e.msg
}

func (e *err) IsServer() bool {
	return e.group == server
}

func (e *err) IsRepository() bool {
	return e.group == repository
}

func (e *err) IsService() bool {
	return e.group == service
}

func (e *err) IsHandler() bool {
	return e.group == handler
}

func (e *err) IsDomain() bool {
	return e.group == domain
}

func (e *err) IsNotFound() bool {
	return e.code == notFoundCode
}

func (e *err) IsInternal() bool {
	return e.code == internalCode
}

func (e *err) IsBadRequest() bool {
	return e.code == badRequestCode
}

func (e *err) IsAlredyExists() bool {
	return e.code == alreadyExistsCode
}

func (e *err) GetGroup() string {
	return e.group
}

func (e *err) GetKind() string {
	return e.kind
}

func (e *err) GetCode() int {
	return e.code
}
