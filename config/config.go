package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	//Token identifying bot
	Token string
	//BotPrefix is the prefix of bot
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

//ReadConfig read in configuration values
func ReadConfig() error {
	fmt.Println("reading from config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
