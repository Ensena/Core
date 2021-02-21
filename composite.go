package core

import (
	"encoding/json"
)

type CombinedResponse struct {
	Ensena interface{}
	Moodle []interface{}
}

func Mixer(b1 *[]byte, b2 *[]byte, name string) []byte {
	var temp1 interface{}
	var temp2 []interface{}
	json.Unmarshal(*b1, &temp1)
	json.Unmarshal(*b2, &temp2)
	temp := CombinedResponse{temp1, temp2}
	b, err := json.Marshal(&temp)
	if err != nil {

		return *b1
	}
	return b
}
