package event

func GetEvent(code string) (Event, error) {
	return FindByCode(code)
}

func CreateEvent(event Event) (Event, error) {
	return Create(event)
}
