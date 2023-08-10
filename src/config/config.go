package config

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

const configFile = "/etc/mon2http/config.yml"

var Port = 4500
var AccessToken = ""
var UpdateInterval = 15 * time.Second
var NumSamples = 3

// LoadConfig todo.
func LoadConfig(cnfPath string) {
	if cnfPath == "" {
		cnfPath = configFile
	}
	var err error
	cnfPath, err = filepath.Abs(cnfPath)
	if err != nil {
		log.Fatalf("config file %s not found", cnfPath)
	}
	var b []byte
	b, err = os.ReadFile(cnfPath)
	if err != nil {
		log.Fatalf("config file %s not found", cnfPath)
	}
	_ = b
}
