package config

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/ini.v1"
)

type SystemConfig struct {
	DbName    string
	EtcdHost  string
	EtcdPort  string
	LoginPage string
}

type ServiceConfig struct {
	Name    string
	Host    string
	Port    string
	WebPort string
	Feature bool
}

type SerivceConfigMap = map[string]*ServiceConfig

type Config struct {
	System  *SystemConfig
	Service SerivceConfigMap
}

var conf *Config
// var once *sync.Once

func GetConfig() *Config {
	return conf
}

func init() {
	conf = &Config{
		Service: make(map[string]*ServiceConfig),
	}
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("Fail to load ini file")
	}
	sections := file.Sections()

	for _, section := range sections {
		name := section.Name()
		if name == "system" {
			conf.System = loadSystemConfig(section)
		} else {
			stringMap := make(map[string]any)

			for _, key := range section.Keys() {
				stringMap[key.Name()] = section.Key(key.Name()).String()
			}
			result := &ServiceConfig{
				Name: name,
			}
			mapstructure.Decode(stringMap, &result)
			conf.Service[name] = result
		}
	}

}

func loadSystemConfig(systemSection *ini.Section) *SystemConfig {
	systemConfig := SystemConfig{}
	tempMap := make(map[string]any)
	for _, key := range systemSection.Keys() {
		tempMap[key.Name()] = key.String()
	}
	mapstructure.Decode(tempMap, &systemConfig)
	return &systemConfig
}
