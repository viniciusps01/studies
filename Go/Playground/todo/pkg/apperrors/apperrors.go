package apperrors

type BadRequestError struct {
	Message string
}

func (e BadRequestError) Error() string {
	return e.Message
}

type InternalServerError struct {
	Message string
}

func (e InternalServerError) Error() string {
	return e.Message
}

type AuthorizationError struct {
	Message string
}

func (e AuthorizationError) Error() string {
	return e.Message
}

type AuthenticationError struct {
	Message string
}

func (e AuthenticationError) Error() string {
	return e.Message
}

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

type ConflictError struct {
	Message string
}

func (e ConflictError) Error() string {
	return e.Message
}

type RateLimitingError struct {
	Message string
}

func (e RateLimitingError) Error() string {
	return e.Message
}

type ExternalServiceError struct {
	Message string
}

func (e ExternalServiceError) Error() string {
	return e.Message
}

type TimeoutError struct {
	Message string
}

func (e TimeoutError) Error() string {
	return e.Message
}
