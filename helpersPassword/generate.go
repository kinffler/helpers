package helpersPassword

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateSecureSecret generates a random 20-character string.
func Generate() (string, error) {
	b := make([]byte, 50)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
