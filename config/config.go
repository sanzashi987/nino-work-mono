package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	DbName          string
	GatewayPort     string
	EtcdHost        string
	EtcdPort        string
	UserServiceName string
	UserServicePort string
	UserServiceHost string
)

func LoadConfig() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("Fail to load ini file")
	}

	loadDbConfig(file)
	loadUserConfig(file)
	loadGateWay(file)
	loadEtcdConfig(file)
}

func loadGateWay(file *ini.File) {
	GatewayPort = file.Section("gateway").Key("Port").String()
}

func loadDbConfig(file *ini.File) {
	DbName = file.Section("sqlite").Key("DbName").String()
}

func loadUserConfig(file *ini.File) {
	userSection := file.Section("user")
	UserServiceName = userSection.Key("ServiceName").String()
	UserServicePort = userSection.Key("Port").String()
	UserServiceHost = userSection.Key("Host").String()
}

func loadEtcdConfig(file *ini.File) {
	etcdSection :=file.Section("etcd")
	EtcdHost = etcdSection.Key("Host").String()
	EtcdPort = etcdSection.Key("Port").String()
}
