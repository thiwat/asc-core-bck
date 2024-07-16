package types

type Session struct {
	Name   string `json:"name"`
	UserId string `json:"user_id"`
	LineId string `json:"line_id"`
}
