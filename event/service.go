package event

func GetEvent(code string) (Event, error) {
	return FindByCode(code)
}

func CreateEvent(event Event) (Event, error) {
	return Create(event)
}

func ListEvent(page int64, pageSize int64, sort string) (ListOutput, error) {
	return List(page, pageSize, sort)
}
