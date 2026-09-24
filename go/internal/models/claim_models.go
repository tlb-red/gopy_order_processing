package claim_models

import "time"

type Claim struct {
	Id        string
	ClientId  string
	Status    string
	Amount    float64
	CreatedAt time.Time
}
