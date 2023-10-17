package cache

import "errors"

var (
	ErrVaultNotFound        = errors.New("vault not found")
	ErrIlkNotFound          = errors.New("ilk not found")
	ErrEraNotFound          = errors.New("era not found")
	ErrAccountNotFound      = errors.New("account not found")
	ErrProtocolDataNotFound = errors.New("protocol data not found")
	ErrAddressedNotFound    = errors.New("addresses not found")
	ErrProxyNotFound        = errors.New("proxy not found")
)
