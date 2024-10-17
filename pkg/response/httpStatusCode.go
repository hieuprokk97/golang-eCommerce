package response

const (
	ErrCodeSuccess             = 200
	ErrCodeBadRequest          = 400
	ErrCodeUnauthorized        = 401
	ErrCodeForbidden           = 403
	ErrCodeNotFound            = 404
	ErrCodeInternalServerError = 500
	ErrCodeParamInvalid        = 422
)

var msg = map[int]string{
	ErrCodeSuccess:             "Success",
	ErrCodeBadRequest:          "Bad Request",
	ErrCodeUnauthorized:        "Unauthorized",
	ErrCodeForbidden:           "Forbidden",
	ErrCodeNotFound:            "Not Found",
	ErrCodeInternalServerError: "Internal Server Error",
	ErrCodeParamInvalid:        "Email is Invalid",
}
