package payloads

type PlayerRequest struct {
	Nickname string `json:"nickname" validate:"required,min=2,max=16"`
	PhoneNo  string `json:"phoneNo" validate:"required,min=11,max=15"`
	Postcode string `json:"postcode" validate:"required,min=5,max=8"`
	City     string `json:"city" validate:"required,min=3,max=50"`
}
type JoinMatchResponse struct {
	MatchId *int `json:"matchId"`
}

type PlayerPaymentRequest struct {
	AmountToPay float32 `json:"amountToPay" validate:"required"`
}
