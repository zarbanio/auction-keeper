package cache

import "errors"

var (
	ErrVaultNotFound     = errors.New("vault not found")
	ErrIlkNotFound       = errors.New("ilk not found")
	ErrEraNotFound       = errors.New("era not found")
	ErrAddressedNotFound = errors.New("addresses not found")
)
