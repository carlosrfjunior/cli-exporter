package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

//LoadFile ...
func (cfg *Config) LoadFile() error {

	filename := "clis.yaml"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Printf("The %s file don't exist!\n", filename)
		return err
	}

	// Open our yamlFile
	file, err := os.Open(filename)

	if err != nil {
		log.Println(err)
	}

	fileRead, _ := ioutil.ReadAll(file)

	err = yaml.Unmarshal([]byte(fileRead), &cfg)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	return nil

}
