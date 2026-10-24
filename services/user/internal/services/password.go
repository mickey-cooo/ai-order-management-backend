package services

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	argonTime    = 1
	argonMemory  = 64 * 1024 // KiB
	argonThreads = 4
	argonKeyLen  = 32
	saltLen      = 16
)

var ErrInvalidHash = errors.New("invalid password hash format")

// preHash is the sha1 step: password -> hex(sha1). Argon2id then hashes the hex string.
func preHash(password string) []byte {
	sum := sha1.Sum([]byte(password))
	return []byte(hex.EncodeToString(sum[:]))
}

// HashPassword returns a PHC-format argon2id hash of sha1(password).
func HashPassword(password string) (string, error) {
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	key := argon2.IDKey(preHash(password), salt, argonTime, argonMemory, argonThreads, argonKeyLen)
	encoded := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, argonMemory, argonTime, argonThreads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(key))
	return encoded, nil
}

// VerifyPassword reports whether password matches an encoded hash from HashPassword.
func VerifyPassword(password, encoded string) (bool, error) {
	parts := strings.Split(encoded, "$")
	if len(parts) != 6 || parts[1] != "argon2id" {
		return false, ErrInvalidHash
	}

	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil || version != argon2.Version {
		return false, ErrInvalidHash
	}

	// Cost settings come from the stored hash, so old hashes still verify after the defaults change.
	var memory, iterations uint32
	var threads uint8
	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &threads); err != nil || iterations < 1 || threads < 1 {
		return false, ErrInvalidHash
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, ErrInvalidHash
	}
	storedKey, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil || len(storedKey) == 0 {
		return false, ErrInvalidHash
	}

	// Key length comes from the stored key, not argonKeyLen.
	newKey := argon2.IDKey(preHash(password), salt, iterations, memory, threads, uint32(len(storedKey)))

	// Constant time, so timing does not leak how many bytes matched.
	return subtle.ConstantTimeCompare(newKey, storedKey) == 1, nil
}

// ComparePassword reports whether password matches hashedPassword. A malformed hash counts as no match.
func ComparePassword(hashedPassword, password string) bool {
	ok, err := VerifyPassword(password, hashedPassword)
	return err == nil && ok
}

// GeneratePassword returns a random URL-safe password of the given length.
func GeneratePassword(length int) (string, error) {
	b := make([]byte, length) // length bytes encode to ~1.33*length chars, so truncating is safe
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b)[:length], nil
}
