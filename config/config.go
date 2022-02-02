package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		DbUri        string `yaml:"uri"`
		DbName       string `yaml:"name"`
		CollSegments string `yaml:"collection_segments"`
	} `yaml:"database"`
	Server struct {
		ServerHost  string `yaml:"host"`
		ServerPort  string `yaml:"port"`
		CronJobTime int    `yaml:"cron_job_interval"`
	} `yaml:"server"`
}

func NewConfig() *Config {

	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var cfg Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}
