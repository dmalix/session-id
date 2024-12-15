package session_id

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func Generate(userId uuid.UUID) (string, error) {

	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	currentTimeMillis := time.Now().UnixNano() / int64(time.Millisecond)
	uniqueData := fmt.Sprintf("%s%s%d", userId, hex.EncodeToString(bytes), currentTimeMillis)
	hsUniqueData := sha256.New()
	hsUniqueData.Write([]byte(uniqueData))
	hashedUniqueDataHex := hex.EncodeToString(hsUniqueData.Sum(nil))

	hashedUniqueData := []byte(hashedUniqueDataHex)

	transformed := make([]byte, 64)

	for i := 0; i < 32; i++ {
		transformed[2*i] = hashedUniqueData[i]
		transformed[2*i+1] = hashedUniqueData[31-i]
	}

	return string(transformed), nil
}

func Check(sessionId string) error {

	if len(sessionId) != 64 {
		return errors.New("the 'session-id' value (" + sessionId + ") is invalid (len64)")
	}

	for i := 0; i < 32; i++ {
		if sessionId[i*2] != sessionId[63-(i*2)] {
			return errors.New("the 'session-id' value (" + sessionId + ") is invalid at position " + string(rune(i*2)))
		}
	}

	return nil
}
