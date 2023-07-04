package errorcode

import (
	code "github.com/marmotedu/sample-code"
)

// 密钥错误
const (
	_ = iota + code.ErrSecretNotFound
	ErrSecretEmpty
)

// 存储桶错误
const (
	_ = iota + code.ErrTokenInvalid
	ErrOssConnectFailed
	ErrOssBucketUploadFailed
)

// HTTP Error
var (
	HttpSuccess                  = NewHttpErr(0, 200, "OK", nil)
	ErrHttpUnknown               = NewHttpErr(code.ErrUnknown, 500, "Internal server error", nil)
	ErrHttpBind                  = NewHttpErr(code.ErrBind, 400, "Error occurred while binding the request body to the struct", nil)
	ErrHttpValidation            = NewHttpErr(code.ErrValidation, 400, "Validation failed", nil)
	ErrHttpTokenInvalid          = NewHttpErr(code.ErrTokenInvalid, 401, "Token invalid", nil)
	ErrHttpDatabase              = NewHttpErr(code.ErrDatabase, 500, "Database error", nil)
	ErrHttpEncrypt               = NewHttpErr(code.ErrEncrypt, 401, "Error occurred while encrypting the user password", nil)
	ErrHttpSignatureInvalid      = NewHttpErr(code.ErrSignatureInvalid, 401, "Signature is invalid", nil)
	ErrHttpExpired               = NewHttpErr(code.ErrExpired, 401, "Token expired", nil)
	ErrHttpInvalidAuthHeader     = NewHttpErr(code.ErrInvalidAuthHeader, 401, "Invalid authorization header", nil)
	ErrHttpMissingHeader         = NewHttpErr(code.ErrMissingHeader, 401, "The `Authorization` header was empty", nil)
	ErrHttporExpired             = NewHttpErr(code.ErrorExpired, 401, "Token expired", nil)
	ErrHttpPasswordIncorrect     = NewHttpErr(code.ErrPasswordIncorrect, 401, "Password was incorrect", nil)
	ErrHttpPermissionDenied      = NewHttpErr(code.ErrPermissionDenied, 403, "Permission denied", nil)
	ErrHttpEncodingFailed        = NewHttpErr(code.ErrEncodingFailed, 500, "Encoding failed due to an error with the data", nil)
	ErrHttpDecodingFailed        = NewHttpErr(code.ErrDecodingFailed, 500, "Decoding failed due to an error with the data", nil)
	ErrHttpInvalidJSON           = NewHttpErr(code.ErrInvalidJSON, 500, "Data is not valid JSON", nil)
	ErrHttpEncodingJSON          = NewHttpErr(code.ErrEncodingJSON, 500, "JSON data could not be encoded", nil)
	ErrHttpDecodingJSON          = NewHttpErr(code.ErrDecodingJSON, 500, "JSON data could not be decoded", nil)
	ErrHttpInvalidYaml           = NewHttpErr(code.ErrInvalidYaml, 500, "Data is not valid Yaml", nil)
	ErrHttpEncodingYaml          = NewHttpErr(code.ErrEncodingYaml, 500, "Yaml data could not be encoded", nil)
	ErrHttpDecodingYaml          = NewHttpErr(code.ErrDecodingYaml, 500, "Yaml data could not be decoded", nil)
	ErrHttpUserNotFound          = NewHttpErr(code.ErrUserNotFound, 404, "User not found", nil)
	ErrHttpUserAlreadyExist      = NewHttpErr(code.ErrUserAlreadyExist, 400, "User already exist", nil)
	ErrHttpReachMaxCount         = NewHttpErr(code.ErrReachMaxCount, 400, "password or username reach the max count", nil)
	ErrHttpSecretEmptyData       = NewHttpErr(ErrSecretEmpty, 400, "密码后用户名不能为空", nil)
	ErrHttpSecretNotFound        = NewHttpErr(code.ErrSecretNotFound, 404, "Secret not found", nil)
	ErrHttpOssConnectFailed      = NewHttpErr(ErrOssConnectFailed, 404, "Oss bucket error", nil)
	ErrHttpOssBucketUploadFailed = NewHttpErr(ErrOssBucketUploadFailed, 404, "Oss bucket upload error", nil)
)

// // Server Error
// var (
// 	Success         = NewErrorCode(0, "OK", nil)
// 	ErrUnknown      = NewErrorCode(code.ErrUnknown, "Internal server error", nil)
// 	ErrBind         = NewErrorCode(code.ErrBind, "Error occurred while binding the request body to the struct", nil)
// 	ErrValidation   = NewErrorCode(code.ErrValidation, "Validation failed", nil)
// 	ErrTokenInvalid = NewErrorCode(code.ErrTokenInvalid, "Token invalid", nil)

// 	ErrDatabase = NewErrorCode(code.ErrDatabase, "Database error", nil)

// 	ErrEncrypt           = NewErrorCode(code.ErrEncrypt, "Error occurred while encrypting the user password", nil)
// 	ErrSignatureInvalid  = NewErrorCode(code.ErrSignatureInvalid, "Signature is invalid", nil)
// 	ErrExpired           = NewErrorCode(code.ErrExpired, "Token expired", nil)
// 	ErrInvalidAuthHeader = NewErrorCode(code.ErrInvalidAuthHeader, "Invalid authorization header", nil)
// 	ErrMissingHeader     = NewErrorCode(code.ErrMissingHeader, "The `Authorization` header was empty", nil)
// 	ErrorExpired         = NewErrorCode(code.ErrorExpired, "Token expired", nil)
// 	ErrPasswordIncorrect = NewErrorCode(code.ErrPasswordIncorrect, "Password was incorrect", nil)
// 	ErrPermissionDenied  = NewErrorCode(code.ErrPermissionDenied, "Permission denied", nil)

// 	ErrEncodingFailed = NewErrorCode(code.ErrEncodingFailed, "Encoding failed due to an error with the data", nil)
// 	ErrDecodingFailed = NewErrorCode(code.ErrDecodingFailed, "Decoding failed due to an error with the data", nil)
// 	ErrInvalidJSON    = NewErrorCode(code.ErrInvalidJSON, "Data is not valid JSON", nil)
// 	ErrEncodingJSON   = NewErrorCode(code.ErrEncodingJSON, "JSON data could not be encoded", nil)
// 	ErrDecodingJSON   = NewErrorCode(code.ErrDecodingJSON, "JSON data could not be decoded", nil)
// 	ErrInvalidYaml    = NewErrorCode(code.ErrInvalidYaml, "Data is not valid Yaml", nil)
// 	ErrEncodingYaml   = NewErrorCode(code.ErrEncodingYaml, "Yaml data could not be encoded", nil)
// 	ErrDecodingYaml   = NewErrorCode(code.ErrDecodingYaml, "Yaml data could not be decoded", nil)

// 	ErrUserNotFound     = NewErrorCode(code.ErrUserNotFound, "User not found", nil)
// 	ErrUserAlreadyExist = NewErrorCode(code.ErrUserAlreadyExist, "User already exist", nil)
// 	ErrReachMaxCount    = NewErrorCode(code.ErrReachMaxCount, "Secret reach the max count", nil)
// 	ErrSecretNotFound   = NewErrorCode(code.ErrSecretNotFound, "Secret not found", nil)
// , nil)
