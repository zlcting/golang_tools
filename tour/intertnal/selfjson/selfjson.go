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

type Gaodecitybylocation struct {
	Status    string `json:"status"`
	Regeocode struct {
		AddressComponent struct {
			City         []interface{} `json:"city"`
			Province     string        `json:"province"`
			Adcode       string        `json:"adcode"`
			District     string        `json:"district"`
			Towncode     string        `json:"towncode"`
			StreetNumber struct {
				Number    string `json:"number"`
				Location  string `json:"location"`
				Direction string `json:"direction"`
				Distance  string `json:"distance"`
				Street    string `json:"street"`
			} `json:"streetNumber"`
			Country       string `json:"country"`
			Township      string `json:"township"`
			BusinessAreas []struct {
				Location string `json:"location"`
				Name     string `json:"name"`
				ID       string `json:"id"`
			} `json:"businessAreas"`
			Building struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"building"`
			Neighborhood struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"neighborhood"`
			Citycode string `json:"citycode"`
		} `json:"addressComponent"`
		FormattedAddress string `json:"formatted_address"`
	} `json:"regeocode"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
}

func Jsonmap(str string) map[string]interface{} {

	dynamic := make(map[string]interface{})

	json.Unmarshal([]byte(str), &dynamic)

	return dynamic
}

func Json2Gaodecitybylocation(str string) Gaodecitybylocation {
	GaodeLoaction := Gaodecitybylocation{}

	json.Unmarshal([]byte(str), &GaodeLoaction)

	return GaodeLoaction
}

func Json2struct(str string) Gaudedistrict {
	district := Gaudedistrict{}

	json.Unmarshal([]byte(str), &district)

	return district
}
