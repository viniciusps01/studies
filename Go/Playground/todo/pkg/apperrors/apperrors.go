package apperrors

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

type DatabaseError struct {
	Message string
}

func (e DatabaseError) Error() string {
	return e.Message
}

type AuthorizationError struct {
	Message string
}

func (e AuthorizationError) Error() string {
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
