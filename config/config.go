package config

import (
	"fmt"
	"sync"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/ini.v1"
)

type SystemConfig struct {
	DbName   string
	EtcdHost string
	EtcdPort string
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
var once *sync.Once

func GetConfig() *Config {
	if conf == nil {
		once.Do(func() {
			conf = &Config{
				Service: make(map[string]*ServiceConfig),
			}
			file, err := ini.Load("./config/config.ini")
			if err != nil {
				fmt.Println("Fail to load ini file")
			}
			loadConfig(file)
		})
	}

	return conf
}

func loadConfig(file *ini.File) {
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
			result := &ServiceConfig{}
			mapstructure.Decode(stringMap, &result)
			conf.Service[name] = result
		}
	}

}

func loadSystemConfig(systemSection *ini.Section) *SystemConfig {
	systemConfig := &SystemConfig{
		EtcdHost: systemSection.Key("EtcdHost").String(),
		EtcdPort: systemSection.Key("EtcdPort").String(),
		DbName:   systemSection.Key("DbName").String(),
	}
	return systemConfig
}
