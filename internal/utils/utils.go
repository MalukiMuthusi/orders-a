package utils

import "errors"

const (
	AppName = "orders"
	Port    = "port"
)

var (
	ErrNoMatchForPhoneNumberRE = errors.New("no match for the phone number REGEXP")
)

const (
	CountryCodesPath   = "COUNTRY_CODES_PATH"
	OrdersPath         = "ORDERS_PATH"
	PostOrdersEndpoint = "ORDERS_POST_ENDPOINT"
	PageSize           = "PAGE_SIZE"
)
