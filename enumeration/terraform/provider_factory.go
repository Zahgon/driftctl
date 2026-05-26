package terraform

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/plugin/discovery"
)

func NewGRPCProvider(meta discovery.PluginMeta) (*plugin.GRPCProvider, error) {
	_ = "STUB: not implemented"
	return nil,

		// Request the RPC terraformProvider so we can get the provider
		// so we can build the actual RPC-implemented provider.
		nil
}
