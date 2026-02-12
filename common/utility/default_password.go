package utility

import (
	"encoding/base64"
	// "fmt"
	"math/rand"

	// sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/scrypt"
)

const (
	SaltBytes = 16
	// OpsLimitModerate from libsodium
	OpsLimit = 3 // MODERATE = 3
	// MemLimitModerate from libsodium (64MB)
	MemLimit = 64 * 1024 * 1024
	// Output length (32 bytes)
	HashLength = 32
)

// func GenerateDefaultPassword(emp sqlc.GetEmployeeByGuidRow) string {
// 	return fmt.Sprint(
// 		emp.DateOfBirth.Time.Format("02012006"),
// 	)
// }

func GenerateSalt() string {
	salt := make([]byte, 16)
	rand.Read(salt)
	return base64.URLEncoding.EncodeToString(salt)
}

func HashPassword(password, salt string) string {
	saltedPassword := []byte(password + salt)
	hashedPassword, _ := scrypt.Key(saltedPassword, []byte(salt), 16384, 8, 1, 32)
	return base64.URLEncoding.EncodeToString(hashedPassword)
}

func GenerateSaltLaravel() ([]byte, error) {
	salt := make([]byte, SaltBytes)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func GenerateSaltLaravelBase64() (string, error) {
	salt, err := GenerateSaltLaravel()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func HashPasswordLaravel(password string, salt []byte) string {
	// Convert password to bytes
	passwordBytes := []byte(password)

	// Perform Argon2id hashing (same algorithm as sodium_crypto_pwhash)
	hashed := argon2.IDKey(
		passwordBytes,
		salt,
		OpsLimit,      // time cost (iterations)
		MemLimit/1024, // memory cost in KiB
		1,             // parallelism
		HashLength,    // output length
	)

	return base64.StdEncoding.EncodeToString(hashed)
}
