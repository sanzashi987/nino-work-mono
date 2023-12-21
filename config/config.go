package config

import (
	"fmt"
	"sync"

	"gopkg.in/ini.v1"
)

type Config struct {
	DbName             string
	GatewayPort        string
	EtcdHost           string
	EtcdPort           string
	UserServiceName    string
	UserServiceHost    string
	UserServicePort    string
	UserServiceWebPort string
}

var config *Config
var once *sync.Once

func InitConfig() *Config {
	if config == nil {
		once.Do(func() {
			file, err := ini.Load("./config/config.ini")
			config = &Config{}
			if err != nil {
				fmt.Println("Fail to load ini file")
			}

			loadDbConfig(file)
			loadUserConfig(file)
			loadGateWay(file)
			loadEtcdConfig(file)
		})
	}

	return config
}

func loadGateWay(file *ini.File) {
	config.GatewayPort = file.Section("gateway").Key("Port").String()
}

func loadDbConfig(file *ini.File) {
	config.DbName = file.Section("sqlite").Key("DbName").String()
}

func loadUserConfig(file *ini.File) {
	userSection := file.Section("user")
	config.UserServiceName = userSection.Key("ServiceName").String()
	config.UserServiceHost = userSection.Key("Host").String()
	config.UserServicePort = userSection.Key("Port").String()
	config.UserServiceWebPort = userSection.Key("WebPort").String()
}

func loadEtcdConfig(file *ini.File) {
	etcdSection := file.Section("etcd")
	config.EtcdHost = etcdSection.Key("Host").String()
	config.EtcdPort = etcdSection.Key("Port").String()
}
