package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts api key from http headers
// Example:
// Authorization: Apikey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no api_key found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "Apikey" {
		return "", errors.New("malformed [0] auth header")
	}

	return vals[1], nil
}
