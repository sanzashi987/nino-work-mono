package service

import (
	"encoding/json"

	"github.com/sanzashi987/nino-work/config"
)

var defatultFeaturesJson string

type Auth struct {
	Read  bool
	Write bool
	Super bool
}

func init() {

	featureMap := make(map[string]Auth)
	serviceMap := config.GetConfig().Service
	for name, service := range serviceMap {
		if service.Feature {
			featureMap[name] = Auth{
				Read:  false,
				Write: false,
				Super: false,
			}
		}
	}

	str, _ := json.Marshal(featureMap)
	defatultFeaturesJson = string(str)

}
