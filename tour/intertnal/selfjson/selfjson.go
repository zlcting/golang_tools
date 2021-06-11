package selfjson

import "encoding/json"

type Gaudedistrict struct {
}

func Jsonmap(str string) map[string]interface{} {

	dynamic := make(map[string]interface{})

	json.Unmarshal([]byte(str), &dynamic)

	return dynamic
}

func Json2struct(str string) Gaudedistrict {
	district := Gaudedistrict{}

	json.Unmarshal([]byte(str), &district)

	return district
}
