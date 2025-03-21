package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken create a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken check if the token is a valid or not
	VerifyToken(token string) (*Payload, error)
}
