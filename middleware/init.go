package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"time"
)

var key1, key2 string

func generateKey(secret string, t time.Time) string {
	data := t.Format("2006-01-02 15:04")
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func generateKeys() {
	for {
		now := time.Now()
		key1 = generateKey(os.Getenv("secretApp"), now)
		key2 = generateKey(os.Getenv("secretApp"), now.Add(-1*time.Minute))
		time.Sleep(59 * time.Second)
	}
}

func init() {
	go generateKeys()
}
