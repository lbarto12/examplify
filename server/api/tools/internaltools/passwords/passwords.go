package passwords

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Errors
var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
	ErrPasswordNotMatch    = errors.New("password does not match encoded hash")
)

type PasswordGenerationOptions struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var numOptionFields int = reflect.TypeOf(PasswordGenerationOptions{}).NumField()

func NewDefaultPasswordGenerationOptions() *PasswordGenerationOptions {
	return &PasswordGenerationOptions{
		memory:      64 * 1024, // 64mb
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
}

func GenerateFromPassword(password string, options *PasswordGenerationOptions) (string, error) {
	if options == nil { // use default if no config provided
		options = NewDefaultPasswordGenerationOptions()
	}
	salt, err := generateRandomBytes(options.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt, options.iterations,
		options.memory,
		options.parallelism,
		options.keyLength,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		options.memory,
		options.iterations,
		options.parallelism,
		b64Salt,
		b64Hash,
	)

	return encodedHash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CompareHashAndPassword(password string, encodedHash string) error {
	result, err := decodeHash(encodedHash)
	if err != nil {
		return err
	}

	passwordHash := argon2.IDKey([]byte(password), result.Salt, result.Options.iterations, result.Options.memory, result.Options.parallelism, result.Options.keyLength)

	if subtle.ConstantTimeCompare(result.Hash, passwordHash) == 1 {
		return nil
	}
	return ErrPasswordNotMatch
}

type DecodedHashResult struct {
	Options *PasswordGenerationOptions
	Salt    []byte
	Hash    []byte
}

func decodeHash(encodedHash string) (result *DecodedHashResult, err error) {
	result = &DecodedHashResult{}

	vals := strings.Split(encodedHash, "$")
	if len(vals) != numOptionFields+1 { // number of fields in options + 1 for version field
		return nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, err
	}
	if version != argon2.Version {
		return nil, ErrIncompatibleVersion
	}

	result.Options = &PasswordGenerationOptions{}
	_, err = fmt.Sscanf(
		vals[3],
		"m=%d,t=%d,p=%d",
		&result.Options.memory,
		&result.Options.iterations,
		&result.Options.parallelism,
	)
	if err != nil {
		return nil, err
	}

	result.Salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, err
	}

	saltLen := len(result.Salt)
	result.Options.saltLength = uint32(saltLen) //nolint:gosec // disable G115

	result.Hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, err
	}

	result.Options.keyLength = uint32(len(result.Hash)) //nolint:gosec // disable G115

	return result, nil
}
