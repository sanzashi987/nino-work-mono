package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	DbName      string
	GatewayPort string
)

func LoadConfig() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("Fail to load ini file")
	}

	LoadDbConfig(file)
}

func LoadGateWay(file *ini.File) {
	GatewayPort = file.Section("Gateway").Key("port").String()
}

func LoadDbConfig(file *ini.File) {
	DbName = file.Section("sqlite").Key("DbName").String()
}
