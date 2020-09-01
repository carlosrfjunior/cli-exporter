package awscli

import (
	"fmt"
	"log"
	"os"

	Config "github.com/carlosrfjunior/cli-exporter/config"
)

// Aws ...
type Aws struct {
	Name string
}

// Start ...
func (a Aws) Start() {

	cfg := Config.Once()

	if err := cfg.LoadFile(); os.IsNotExist(err) {
		log.Printf("The clis.yaml file not found!\n")
	}

	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", cfg)
	fmt.Printf("%v\n", cfg.GetProviderService("ses-email-verify-count"))
}
