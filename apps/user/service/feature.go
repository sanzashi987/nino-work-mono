package service

import (
	"encoding/json"

	"github.com/sanzashi987/nino-work/config"
)

var defatultFeaturesJson string

func init() {

	featureMap := make(map[string]bool)
	serviceMap := config.GetConfig().Service
	for name, service := range serviceMap {
		if service.Feature {
			featureMap[name] = false
		}
	}

	str, _ := json.Marshal(featureMap)
	defatultFeaturesJson = string(str)

}
