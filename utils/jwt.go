package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)


func Decode(jwtToken string, payload interface{}) (error) {
	token := strings.Split(jwtToken, ".")
	// check if the jwtToken token contains
	// header, payload and token
	if len(token) != 3 {
		splitErr := errors.New("invalid token: token should contain header, payload and secret")
		return splitErr
	}
	// decode payload
	decodedPayload, PayloadErr := base64Decode(token[1])
	if PayloadErr != nil {
		return fmt.Errorf("invalid payload: %s", PayloadErr.Error())
	}

	// parses payload from string to a struct
	ParseErr := json.Unmarshal([]byte(decodedPayload), &payload)
	if ParseErr != nil {
		return fmt.Errorf("invalid payload: %s", ParseErr.Error())
	}

	return nil
}

// Base64Encode takes in a string and returns a base 64 encoded string
//func Base64Encode(src string) string {
//	return strings.
//		TrimRight(base64.URLEncoding.
//			EncodeToString([]byte(src)), "=")
//}

// Base64Encode takes in a base 64 encoded string and returns the //actual string or an error of it fails to decode the string
func base64Decode(src string) (string, error) {
	if l:= len(src) % 4; l > 0 {
		src += strings.Repeat("=", 4-l)
	}
	decoded, err := base64.URLEncoding.DecodeString(src)
	if err != nil {
		errMsg := fmt.Errorf("decoding Error %s", err)
		return "", errMsg
	}
	return string(decoded), nil
}