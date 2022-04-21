package payloads

type SquadRequest struct {
	Name string `json:"name" validate:"required,min=2,max=30"`
	City string `json:"city" validate:"required,min=3,max=50"`
}
