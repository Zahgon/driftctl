package aws

import (
	"github.com/snyk/driftctl/enumeration/alerter"
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type wrongArnTopicAlert struct {
	arn      string
	endpoint *string
}

func NewWrongArnTopicAlert(arn string, endpoint *string) *wrongArnTopicAlert {
	_ = "STUB: not implemented"
	return nil
}

func (p *wrongArnTopicAlert) Message() string { _ = "STUB: not implemented"; return "" }

func (p *wrongArnTopicAlert) ShouldIgnoreResource() bool { _ = "STUB: not implemented"; return false }

func (p *wrongArnTopicAlert) Resource() *resource.Resource { _ = "STUB: not implemented"; return nil }

type SNSTopicSubscriptionEnumerator struct {
	repository repository.SNSRepository
	factory    resource.ResourceFactory
	alerter    alerter.AlerterInterface
}

func NewSNSTopicSubscriptionEnumerator(
	repo repository.SNSRepository,
	factory resource.ResourceFactory,
	alerter alerter.AlerterInterface,
) *SNSTopicSubscriptionEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *SNSTopicSubscriptionEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *SNSTopicSubscriptionEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
