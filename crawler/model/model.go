package model

import "encoding/json"

type Profile struct {
	Name          string
	City          string
	Gender        string
	Age           string
	Height        string
	Weight        string
	Income        string
	Marriage      string
	Education     string
	Nation        string
	House         string
	Car           string
	Residence     string
	Constellation string
}

func GetProfileFromJson(o interface{}) (Profile, error) {
	var profile Profile

	bytes, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(bytes, &profile)

	return profile, nil
}
