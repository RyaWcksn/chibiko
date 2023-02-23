package encryptions

import (
	"encoding/base64"
	"strconv"
)

// Encode number to base62.
func Encode(n int64) string {
	data := strconv.Itoa(int(n))
	return base64.StdEncoding.EncodeToString([]byte(data))

}

// Decode converts a base62 token to int.
func Decode(key string) (int64, error) {
	decoded, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return 0, err
	}
	ret, err := strconv.Atoi(string(decoded))
	if err != nil {
		return 0, err
	}

	return int64(ret), nil

}
