package lock

type ProviderBlock struct {
	Address     string   `hcl:"address,label"`
	Version     string   `hcl:"version,attr"`
	Constraints string   `hcl:"constraints,optional"`
	Hashes      []string `hcl:"hashes,optional"`
}

// ProviderAddress encapsulates a single provider type. In the future this will be
// extended to include additional fields including Namespace and SourceHost
type ProviderAddress struct {
	Type      string
	Namespace string
	Hostname  string
}

func (p *ProviderAddress) String() string { _ = "STUB: not implemented"; return "" }

type Lockfile struct {
	Providers []ProviderBlock `hcl:"provider,block"`
}

func (l *Lockfile) GetProviderByAddress(addr *ProviderAddress) *ProviderBlock {
	_ = "STUB: not implemented"
	return nil
}

func ReadLocksFromFile(filename string) (*Lockfile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
