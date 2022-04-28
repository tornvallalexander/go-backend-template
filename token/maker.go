package token

import "time"

// Maker is interface for managing tokens
type Maker interface {
	// CreateToken creates a token with specified username and duration time
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
