package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

// createAt time.Now().Unix()
func GenerateSlug(title string, createAt int64) (string, error) {

	hashInput := fmt.Sprintf("%s%d", title, createAt)
	hash := md5.Sum([]byte(hashInput))
	hashStr := hex.EncodeToString(hash[:])[:12]

	num, err := strconv.ParseInt(hashStr, 16, 64)
	if err != nil {
		return "", err
	}

	slug := strconv.FormatInt(num, 32)

	return slug, nil
}
