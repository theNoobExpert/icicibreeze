package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/go-playground/validator/v10"
)

func CalculateChecksum(payload, secret_key string) (string, string) {
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05") + ".000Z"
	data := timestamp + payload + secret_key
	hash := sha256.Sum256([]byte(data))

	checksum := hex.EncodeToString(hash[:])
	return checksum, timestamp
}

func AsPtr(s string) *string {
	return &s
}

var Validate = validator.New()
