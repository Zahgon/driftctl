package terraform

import (
	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform/plugin/discovery"
)

func ClientConfig(m discovery.PluginMeta) *plugin.ClientConfig {
	_ = "STUB: not implemented"
	return nil
}

// Client returns a plugin client for the plugin described by the given metadata.
func Client(m discovery.PluginMeta) *plugin.Client { _ = "STUB: not implemented"; return nil }
