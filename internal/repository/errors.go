package repository

import "errors"

var (
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrRecordNotFound    = errors.New("record not found")
)
