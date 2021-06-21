package selfjson

import "encoding/json"

// Gaudedistrict https://mholt.github.io/json-to-go/ 自动生成json对应的struct
type Gaudedistrict struct {
	Status     string `json:"status"`
	Info       string `json:"info"`
	Infocode   string `json:"infocode"`
	Count      string `json:"count"`
	Suggestion struct {
		Keywords []interface{} `json:"keywords"`
		Cities   []interface{} `json:"cities"`
	} `json:"suggestion"`
	Districts []struct {
		Citycode  string        `json:"citycode"`
		Adcode    string        `json:"adcode"`
		Name      string        `json:"name"`
		Polyline  string        `json:"polyline"`
		Center    string        `json:"center"`
		Level     string        `json:"level"`
		Districts []interface{} `json:"districts"`
	} `json:"districts"`
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
