package service

import (
	"encoding/json"
	"sync"

	"github.com/cza14h/nino-work/config"
)

var defatultFeaturesJson string
var onceFeature *sync.Once

func init() {

	once.Do(func() {
		featureMap := make(map[string]bool)
		serviceMap := config.GetConfig().Service
		for name, service := range serviceMap {
			if service.Feature {
				featureMap[name] = false
			}
		}

		str, _ := json.Marshal(featureMap)
		defatultFeaturesJson = string(str)

	})

}
