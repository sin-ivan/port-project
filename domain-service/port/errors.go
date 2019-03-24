package port

import "errors"

var (
	INTERNAL_SERVER_ERROR = errors.New("Internal Server Error")
	NOT_FOUND_ERROR       = errors.New("Requested Item is not found")
	ITEM_EXIST_ERROR      = errors.New("Item already exist")
)
