package vsphere

import (
	"github.com/cloudfoundry/bosh-bootloader/bosh"
	"github.com/cloudfoundry/bosh-bootloader/storage"
)

type InputGenerator struct {
}

func NewInputGenerator() InputGenerator {
	return InputGenerator{}
}

func (i InputGenerator) Generate(state storage.State) (map[string]interface{}, error) {
	cidr := state.VSphere.Subnet
	parsedCidr, _ := bosh.ParseCIDRBlock(cidr)
	return map[string]interface{}{
		"vsphere_subnet": cidr,
		"jumpbox_ip":     parsedCidr.GetNthIP(10).String(),
	}, nil
}
