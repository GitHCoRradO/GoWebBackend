package request

import "time"

type Bankcard struct {
	Number      string    `json:"name"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}