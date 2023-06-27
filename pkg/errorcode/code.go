package errorcode

import (
	code "github.com/marmotedu/sample-code"
)

// HTTP Error
var (
	HttpSuccess              = NewHttpErr(0, 200, "OK")
	ErrHttpUnknown           = NewHttpErr(code.ErrUnknown, 500, "Internal server error")
	ErrHttpBind              = NewHttpErr(code.ErrBind, 400, "Error occurred while binding the request body to the struct")
	ErrHttpValidation        = NewHttpErr(code.ErrValidation, 400, "Validation failed")
	ErrHttpTokenInvalid      = NewHttpErr(code.ErrTokenInvalid, 401, "Token invalid")
	ErrHttpDatabase          = NewHttpErr(code.ErrDatabase, 500, "Database error")
	ErrHttpEncrypt           = NewHttpErr(code.ErrEncrypt, 401, "Error occurred while encrypting the user password")
	ErrHttpSignatureInvalid  = NewHttpErr(code.ErrSignatureInvalid, 401, "Signature is invalid")
	ErrHttpExpired           = NewHttpErr(code.ErrExpired, 401, "Token expired")
	ErrHttpInvalidAuthHeader = NewHttpErr(code.ErrInvalidAuthHeader, 401, "Invalid authorization header")
	ErrHttpMissingHeader     = NewHttpErr(code.ErrMissingHeader, 401, "The `Authorization` header was empty")
	ErrHttporExpired         = NewHttpErr(code.ErrorExpired, 401, "Token expired")
	ErrHttpPasswordIncorrect = NewHttpErr(code.ErrPasswordIncorrect, 401, "Password was incorrect")
	ErrHttpPermissionDenied  = NewHttpErr(code.ErrPermissionDenied, 403, "Permission denied")
	ErrHttpEncodingFailed    = NewHttpErr(code.ErrEncodingFailed, 500, "Encoding failed due to an error with the data")
	ErrHttpDecodingFailed    = NewHttpErr(code.ErrDecodingFailed, 500, "Decoding failed due to an error with the data")
	ErrHttpInvalidJSON       = NewHttpErr(code.ErrInvalidJSON, 500, "Data is not valid JSON")
	ErrHttpEncodingJSON      = NewHttpErr(code.ErrEncodingJSON, 500, "JSON data could not be encoded")
	ErrHttpDecodingJSON      = NewHttpErr(code.ErrDecodingJSON, 500, "JSON data could not be decoded")
	ErrHttpInvalidYaml       = NewHttpErr(code.ErrInvalidYaml, 500, "Data is not valid Yaml")
	ErrHttpEncodingYaml      = NewHttpErr(code.ErrEncodingYaml, 500, "Yaml data could not be encoded")
	ErrHttpDecodingYaml      = NewHttpErr(code.ErrDecodingYaml, 500, "Yaml data could not be decoded")
	ErrHttpUserNotFound      = NewHttpErr(code.ErrUserNotFound, 404, "User not found")
	ErrHttpUserAlreadyExist  = NewHttpErr(code.ErrUserAlreadyExist, 400, "User already exist")
	ErrHttpReachMaxCount     = NewHttpErr(code.ErrReachMaxCount, 400, "Secret reach the max count")
	ErrHttpSecretNotFound    = NewHttpErr(code.ErrSecretNotFound, 404, "Secret not found")
)

// Server Error
var (
	Success         = NewErrorCode(0, "OK")
	ErrUnknown      = NewErrorCode(code.ErrUnknown, "Internal server error")
	ErrBind         = NewErrorCode(code.ErrBind, "Error occurred while binding the request body to the struct")
	ErrValidation   = NewErrorCode(code.ErrValidation, "Validation failed")
	ErrTokenInvalid = NewErrorCode(code.ErrTokenInvalid, "Token invalid")

	ErrDatabase = NewErrorCode(code.ErrDatabase, "Database error")

	ErrEncrypt           = NewErrorCode(code.ErrEncrypt, "Error occurred while encrypting the user password")
	ErrSignatureInvalid  = NewErrorCode(code.ErrSignatureInvalid, "Signature is invalid")
	ErrExpired           = NewErrorCode(code.ErrExpired, "Token expired")
	ErrInvalidAuthHeader = NewErrorCode(code.ErrInvalidAuthHeader, "Invalid authorization header")
	ErrMissingHeader     = NewErrorCode(code.ErrMissingHeader, "The `Authorization` header was empty")
	ErrorExpired         = NewErrorCode(code.ErrorExpired, "Token expired")
	ErrPasswordIncorrect = NewErrorCode(code.ErrPasswordIncorrect, "Password was incorrect")
	ErrPermissionDenied  = NewErrorCode(code.ErrPermissionDenied, "Permission denied")

	ErrEncodingFailed = NewErrorCode(code.ErrEncodingFailed, "Encoding failed due to an error with the data")
	ErrDecodingFailed = NewErrorCode(code.ErrDecodingFailed, "Decoding failed due to an error with the data")
	ErrInvalidJSON    = NewErrorCode(code.ErrInvalidJSON, "Data is not valid JSON")
	ErrEncodingJSON   = NewErrorCode(code.ErrEncodingJSON, "JSON data could not be encoded")
	ErrDecodingJSON   = NewErrorCode(code.ErrDecodingJSON, "JSON data could not be decoded")
	ErrInvalidYaml    = NewErrorCode(code.ErrInvalidYaml, "Data is not valid Yaml")
	ErrEncodingYaml   = NewErrorCode(code.ErrEncodingYaml, "Yaml data could not be encoded")
	ErrDecodingYaml   = NewErrorCode(code.ErrDecodingYaml, "Yaml data could not be decoded")

	ErrUserNotFound     = NewErrorCode(code.ErrUserNotFound, "User not found")
	ErrUserAlreadyExist = NewErrorCode(code.ErrUserAlreadyExist, "User already exist")
	ErrReachMaxCount    = NewErrorCode(code.ErrReachMaxCount, "Secret reach the max count")
	ErrSecretNotFound   = NewErrorCode(code.ErrSecretNotFound, "Secret not found")
)
