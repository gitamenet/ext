package conf

import (
	"github.com/gitamenet/v2ray-core/proxy/dns"
	"github.com/golang/protobuf/proto"
)

type DnsOutboundConfig struct{}

func (c *DnsOutboundConfig) Build() (proto.Message, error) {
	return new(dns.Config), nil
}
