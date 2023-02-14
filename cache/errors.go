package cache

import "errors"

var (
	ErrVaultNotFound = errors.New("vault not found")
	ErrIlkNotFound   = errors.New("ilk not found")
)
