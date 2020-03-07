package conf

import (
	"strings"

	"github.com/gitamenet/v2ray-core/app/commander"
	loggerservice "github.com/gitamenet/v2ray-core/app/log/command"
	handlerservice "github.com/gitamenet/v2ray-core/app/proxyman/command"
	statsservice "github.com/gitamenet/v2ray-core/app/stats/command"
	"github.com/gitamenet/v2ray-core/common/serial"
)

type ApiConfig struct {
	Tag      string   `json:"tag"`
	Services []string `json:"services"`
}

func (c *ApiConfig) Build() (*commander.Config, error) {
	if len(c.Tag) == 0 {
		return nil, newError("Api tag can't be empty.")
	}

	services := make([]*serial.TypedMessage, 0, 16)
	for _, s := range c.Services {
		switch strings.ToLower(s) {
		case "handlerservice":
			services = append(services, serial.ToTypedMessage(&handlerservice.Config{}))
		case "loggerservice":
			services = append(services, serial.ToTypedMessage(&loggerservice.Config{}))
		case "statsservice":
			services = append(services, serial.ToTypedMessage(&statsservice.Config{}))
		}
	}

	return &commander.Config{
		Tag:     c.Tag,
		Service: services,
	}, nil
}
