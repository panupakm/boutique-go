package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

var encodedHashFmt = "$argon2id$v=%d$m=%d,t=%d,p=%d$%s %s"

func HashPassword(password string) (string, error) {
	argon2Time := uint32(1)
	argon2Memory := uint32(64 * 1024)
	argon2Threads := uint8(4)
	argon2KeyLen := uint32(32)

	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, argon2Time, argon2Memory, argon2Threads, argon2KeyLen)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf(encodedHashFmt, argon2.Version, argon2Memory, argon2Time, argon2Threads, b64Salt, b64Hash), nil
}

func CheckPasswordHash(password, encodedHash string) (bool, error) {
	var version int
	var memory, times uint32
	var threads uint8
	var b64Salt, b64Hash string

	_, err := fmt.Sscanf(encodedHash, encodedHashFmt, &version, &memory, &times, &threads, &b64Salt, &b64Hash)

	fmt.Println(b64Salt)
	fmt.Println(b64Hash)
	salt, err := base64.RawStdEncoding.DecodeString(b64Salt)
	if err != nil {
		return false, err
	}

	hash, err := base64.RawStdEncoding.DecodeString(b64Hash)
	if err != nil {
		return false, err
	}
	checkHash := argon2.IDKey([]byte(password), salt, times, memory, threads, uint32(len(hash)))

	for idx, b := range hash {
		if checkHash[idx] != b {
			return false, nil
		}
	}

	return true, nil
}
