package token

import "time"

// Maker adalah interface untuk mengelola token
type Maker interface {
	// CreateToken, membuat token baru untuk spesifik username dan duratsi yang valid
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken, mengecek apakah token valid atau tidak
	VerifyToken(token string) (*Payload, error)
}
