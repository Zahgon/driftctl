package terraform

const (
	AWS    string = "aws"
	GITHUB string = "github"
	GOOGLE string = "google"
	AZURE  string = "azurerm"
)

type ProviderLibrary struct {
	providers map[string]TerraformProvider
}

func NewProviderLibrary() *ProviderLibrary { _ = "STUB: not implemented"; return nil }

func (p *ProviderLibrary) AddProvider(name string, provider TerraformProvider) {
	_ = "STUB: not implemented"
	return
}

func (p *ProviderLibrary) Provider(name string) TerraformProvider {
	_ = "STUB: not implemented"
	return *new(TerraformProvider)
}

func (p *ProviderLibrary) Cleanup() { _ = "STUB: not implemented"; return }
