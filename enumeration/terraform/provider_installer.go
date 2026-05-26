package terraform

type HomeDirInterface interface {
	Dir() (string, error)
}

type ProviderInstaller struct {
	downloader ProviderDownloaderInterface
	config     ProviderConfig
	homeDir    string
}

func NewProviderInstaller(config ProviderConfig) (*ProviderInstaller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *ProviderInstaller) Install() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (p ProviderInstaller) getProviderDirectory() string { _ = "STUB: not implemented"; return "" }

// Handle postfixes in binary names
func (p *ProviderInstaller) getBinaryPath() string { _ = "STUB: not implemented"; return "" }
