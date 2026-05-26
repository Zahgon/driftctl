package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"
	"github.com/snyk/driftctl/enumeration"
	"github.com/snyk/driftctl/enumeration/remote/terraform"
)

type awsConfig struct {
	AccessKey     string
	SecretKey     string
	CredsFilename string
	Profile       string
	Token         string
	Region        string `cty:"region"`
	MaxRetries    int

	AssumeRoleARN         string
	AssumeRoleExternalID  string
	AssumeRoleSessionName string
	AssumeRolePolicy      string

	AllowedAccountIds   []string
	ForbiddenAccountIds []string

	Endpoints        map[string]string
	IgnoreTagsConfig map[string]string
	Insecure         bool

	SkipCredsValidation     bool `cty:"skip_credentials_validation"`
	SkipGetEC2Platforms     bool
	SkipRegionValidation    bool
	SkipRequestingAccountId bool `cty:"skip_requesting_account_id"`
	SkipMetadataApiCheck    bool
	S3ForcePathStyle        bool
}

type AWSTerraformProvider struct {
	*terraform.TerraformProvider
	session   *session.Session
	name      string
	version   string
	accountId string
}

func NewAWSTerraformProvider(version string, progress enumeration.ProgressCounter, configDir string) (*AWSTerraformProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Those two parameters are used to make sure that the credentials are not validated when calling
// Configure(). Credentials validation is now handled directly in driftctl

// TODO make this configurable

func (a *AWSTerraformProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (p *AWSTerraformProvider) Version() string { _ = "STUB: not implemented"; return "" }

var AWSCredentialsNotFoundError = errors.New("Could not find a way to authenticate on AWS!\n" +
	"Please refer to AWS documentation: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html")

func (p *AWSTerraformProvider) CheckCredentialsExist() error { _ = "STUB: not implemented"; return nil }

// This call is to make sure that the credentials are valid
// A more complex logic exist in terraform provider, but it's probably not worth to implement it
// https://github.com/hashicorp/terraform-provider-aws/blob/e3959651092864925045a6044961a73137095798/aws/auth_helpers.go#L111
