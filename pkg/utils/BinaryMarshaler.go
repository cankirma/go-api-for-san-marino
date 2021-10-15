package utils

import "encoding/json"

func MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(data)
	return bytes, err
}