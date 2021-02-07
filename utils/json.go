package utils

import "encoding/json"

func StructToJson(value interface{}) (res string, err error) {
	bytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func JsonToStruct(data string, value interface{}) error {
	return json.Unmarshal([]byte(data), value)
}
