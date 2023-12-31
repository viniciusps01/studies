package apperrors

import "net/http"

type HttpError struct {
	Message string
	Status  int
}

func (e HttpError) Error() string {
	return e.Message
}

func HttpErrorFrom(err error) *HttpError {
	switch err.(type) {
	case BadRequestError:
		return &HttpError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}

	case InternalServerError:
		return &HttpError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}

	case AuthenticationError:
		return &HttpError{
			Message: err.Error(),
			Status:  http.StatusUnauthorized,
		}

	case AuthorizationError:
		return &HttpError{
			Message: err.Error(),
			Status:  http.StatusForbidden,
		}

	case NotFoundError:
		return &HttpError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}

	case ConflictError:
		return &HttpError{
			Message: err.Error(),
			Status:  http.StatusConflict,
		}

	case RateLimitingError:
		return &HttpError{
			Message: err.Error(),
			Status:  http.StatusTooManyRequests,
		}

	case ExternalServiceError:
		return &HttpError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}

	case TimeoutError:
		return &HttpError{
			Message: err.Error(),
			Status:  http.StatusRequestTimeout,
		}

	default:
		return &HttpError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
}
