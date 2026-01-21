package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateLicenseKey(domain string) (string, error) {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomPart := hex.EncodeToString(randomBytes)[:16]

	licenseKey := fmt.Sprintf("%s-%s", domain, randomPart)

	return licenseKey, nil
}
