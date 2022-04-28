package payloads

type PitchRequest struct {
	Name       string  `json:"name" validate:"required,min=3,max=100"`
	MaxPlayers int     `json:"maxPlayers" validate:"required,max=22"`
	Cost       float32 `json:"cost" validate:"required"`
}
