package user

type LoginInput struct {
	LineToken string `json:"line_token"`
}

type LoginOutput struct {
	Profile User   `json:"profile"`
	Token   string `json:"token"`
}
