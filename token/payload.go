package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var InvalidTokenError = errors.New("jwt is invalid")
var ExpiredJwtError = errors.New("jwt has expired")

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (payload *Payload, err error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload = &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ExpiredJwtError
	}
	return nil
}
