package utils

import "encoding/json"

func ReMarshal(m, result interface{}) error {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, result)
	if err != nil {
		return err
	}
	return nil
}
