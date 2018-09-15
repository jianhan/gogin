package google

import (
	"github.com/jianhan/gogin/config"
	"googlemaps.github.io/maps"
)

func MapClient() (c *maps.Client, err error) {
	if c, err = maps.NewClient(maps.WithAPIKey(config.GetEnvs().GoogleMapAPIKey)); err != nil {
		return
	}

	return
}
