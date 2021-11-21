package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ConfigService struct {
	Skins       []string `json:"skins"`
	Backgrounds []string `json:"backgrounds"`
	Mouth      []string `json:"mouth"`
	Eyes       []string `json:"eyes"`
	Tracks			 []string `json:"tracks"`
	Vortex		 []string `json:"vortex"`
	Necklaces	 []string `json:"necklaces"`
}

func NewConfigService(configPath string) *ConfigService {
	jsonFile, err := os.Open(configPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var service ConfigService

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &service)

	return &service
}
