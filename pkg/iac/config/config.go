package config

type SupplierConfig struct {
	Key     string
	Backend string
	Path    string
}

func (c *SupplierConfig) String() string { _ = "STUB: not implemented"; return "" }
