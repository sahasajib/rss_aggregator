package auth

import (
	"errors"
	"net/http"
	"strings"
)

//GetAPIKey extects an API Key from
//the headers of an HTTP request
//Example:
//Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error){
	val := headers.Get("Authorization")
	if val == ""{
		return "", errors.New("No authentication info found")
	}

	vals := strings.Fields(val)
	if len(vals) < 2 || vals[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return vals[1], nil
}