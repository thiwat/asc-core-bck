package event

import (
	"asc-core/types"
)

func GetEvent(code string) (Event, error) {
	return FindByCode(code)
}

func CreateEvent(event Event) (Event, error) {
	return Create(event)
}

func ListEvent(page int64, pageSize int64, sort string) (types.ListOutput[Event], error) {
	return List(page, pageSize, sort)
}