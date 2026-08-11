package ai

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	ErrConfigEncryptionKeyUnavailable = errors.New("AI configuration encryption key is unavailable")
	ErrConfigDecryptionFailed         = errors.New("AI configuration decryption failed")
)

// SecretCipher keeps provider credentials encrypted while they are at rest.
type SecretCipher interface {
	Encrypt(plaintext string) (string, error)
	Decrypt(ciphertext string) (string, error)
}

type aesGCMCipher struct {
	aead cipher.AEAD
}

// NewConfigCipher reads a base64-encoded, 32-byte AES-256 key directly from
// the deployment environment. There is intentionally no generated default.
func NewConfigCipher() (SecretCipher, error) {
	rawKey, ok := os.LookupEnv("AI_CONFIG_ENCRYPTION_KEY")
	if !ok || strings.TrimSpace(rawKey) == "" {
		return nil, ErrConfigEncryptionKeyUnavailable
	}
	return NewAESGCMCipher(rawKey)
}

// NewAESGCMCipher constructs a cipher from a base64-encoded AES-256 key.
// It is exported to make the encryption boundary testable without process
// configuration.
func NewAESGCMCipher(encodedKey string) (SecretCipher, error) {
	key, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil || len(key) != 32 {
		return nil, ErrConfigEncryptionKeyUnavailable
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("create AES cipher: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("create AES-GCM cipher: %w", err)
	}
	return &aesGCMCipher{aead: aead}, nil
}

func (c *aesGCMCipher) Encrypt(plaintext string) (string, error) {
	nonce := make([]byte, c.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("generate nonce: %w", err)
	}
	ciphertext := c.aead.Seal(nil, nonce, []byte(plaintext), nil)
	return base64.RawStdEncoding.EncodeToString(append(nonce, ciphertext...)), nil
}

func (c *aesGCMCipher) Decrypt(encodedCiphertext string) (string, error) {
	payload, err := base64.RawStdEncoding.DecodeString(encodedCiphertext)
	if err != nil || len(payload) < c.aead.NonceSize() {
		return "", fmt.Errorf("%w: invalid ciphertext", ErrConfigDecryptionFailed)
	}
	nonce := payload[:c.aead.NonceSize()]
	ciphertext := payload[c.aead.NonceSize():]
	plaintext, err := c.aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("%w: authentication failed", ErrConfigDecryptionFailed)
	}
	return string(plaintext), nil
}
