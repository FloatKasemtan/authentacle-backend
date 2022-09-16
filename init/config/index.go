package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var C = new(config)

func init() {
	fmt.Println("Config is initializing...")
	// Load conf to struct
	yml, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error to read config file")
	}
	err = yaml.Unmarshal(yml, C)
	if err != nil {
		log.Fatal("Error to parse config file")
	}
}
