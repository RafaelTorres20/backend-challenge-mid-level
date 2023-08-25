package domain

import "errors"

var (
	ErrAssetNotFound       = errors.New("asset not found")
	ErrInternalServerError = errors.New("internal server error")
	ErrUserNotFound        = errors.New("user not found")
	ErrInvalidAsset        = errors.New("invalid asset")
)
