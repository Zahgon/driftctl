package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// SQS queues from AWS have a weird behaviour when we fetch them.
// By default they have a Policy attached with only an ID
// "arn:aws:sqs:eu-west-3:XXXXXXXXXXXX:foobar/SQSDefaultPolicy" but on fetch
// the SDK return an empty policy (e.g. policy = "").
// We need to ignore those policy from unmanaged resources if they are not managed
// by IaC.
type AwsDefaultSQSQueuePolicy struct{}

func NewAwsDefaultSQSQueuePolicy() AwsDefaultSQSQueuePolicy {
	_ = "STUB: not implemented"
	return *new(AwsDefaultSQSQueuePolicy)
}

func (m AwsDefaultSQSQueuePolicy) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than sqs_queue_policy

// Ignore all non-default queue policy

// Check if queue policy is managed by IaC

// Include resource if it's managed in IaC

// Else, resource is not added to newRemoteResources slice so it will be ignored
