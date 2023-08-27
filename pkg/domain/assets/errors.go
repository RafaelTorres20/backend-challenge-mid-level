package assets

import "errors"

var (
	ErrAssetNotFound       = errors.New("asset not found")
	ErrInternalServerError = errors.New("internal server error")
	ErrUserNotFound        = errors.New("user not found")
	ErrInvalidAsset        = errors.New("invalid asset")
	ErrUserIDNotFound      = errors.New("user id not found")
	ErrBadRequest          = errors.New("bad request")
)
