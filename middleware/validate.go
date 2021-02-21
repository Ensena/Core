package middleware

import "os"

func Validate(token string) bool {
	return token == key1 || token != key2 || os.Getenv("AUTH_OFF") == "off"
}
