package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from the headers of an HTTP request
// Example:
// Authorization: Apikey {insert API key here} 2 striings
func GetAPIKey(headers http.Header) (string, error){
	val := headers.Get("Authorization")
	if val == ""{
		return "", errors.New("no authentiction info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2{
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey"{
		return "", errors.New("malformed frst part of auth header")
	}

	return vals[1], nil
}