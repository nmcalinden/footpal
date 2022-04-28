package views

type Players struct {
	Pagination Pagination `json:"pagination"`
	Data       []Player   `json:"data"`
}

type Player struct {
	PlayerId *int    `json:"playerId"`
	Name     string  `json:"name"`
	Nickname *string `json:"nickname"`
	City     string  `json:"city" `
}
