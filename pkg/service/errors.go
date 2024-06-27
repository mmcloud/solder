package service

import "errors"

var (
	ErrNotFound               = errors.New("not found")
	ErrServiceConfigNotFound  = errors.New("service config not found")
	ErrUnsupportedServiceKind = errors.New("service kind not found")
)
