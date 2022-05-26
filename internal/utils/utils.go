package utils

import "errors"

const (
	AppName = "orders"
	Port    = "port"
)

var (
	ErrNoMatchForPhoneNumberRE = errors.New("no match for the phone number REGEXP")
)
