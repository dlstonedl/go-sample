package model

import "encoding/json"

type Profile struct {
	Name          string
	City          string
	Gender        string
	Age           int
	Height        int
	Weight        int
	Income        string
	Marriage      string
	Education     string
	Nation        string
	House         string
	Car           string
	Residence     string
	Constellation string
}

func GetProfileFromJson(obj interface{}) (Profile, error) {
	var profile Profile

	bytes, err := json.Marshal(obj)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(bytes, &profile)

	return profile, nil
}
