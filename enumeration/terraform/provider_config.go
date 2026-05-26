package terraform

type ProviderConfig struct {
	Key       string
	Version   string
	ConfigDir string
}

func (c *ProviderConfig) GetDownloadUrl() string { _ = "STUB: not implemented"; return "" }

func (c *ProviderConfig) GetBinaryName() string { _ = "STUB: not implemented"; return "" }
