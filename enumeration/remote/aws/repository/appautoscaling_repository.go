package repository

import (
	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling/applicationautoscalingiface"
)

type AppAutoScalingRepository interface {
	ServiceNamespaceValues() []string
	DescribeScalableTargets(string) ([]*applicationautoscaling.ScalableTarget, error)
	DescribeScalingPolicies(string) ([]*applicationautoscaling.ScalingPolicy, error)
	DescribeScheduledActions(string) ([]*applicationautoscaling.ScheduledAction, error)
}

type appAutoScalingRepository struct {
	client applicationautoscalingiface.ApplicationAutoScalingAPI
	cache  cache.Cache
}

func NewAppAutoScalingRepository(session *session.Session, c cache.Cache) *appAutoScalingRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *appAutoScalingRepository) ServiceNamespaceValues() []string {
	_ = "STUB: not implemented"
	return nil
}

func (r *appAutoScalingRepository) DescribeScalableTargets(namespace string) ([]*applicationautoscaling.ScalableTarget, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *appAutoScalingRepository) DescribeScalingPolicies(namespace string) ([]*applicationautoscaling.ScalingPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *appAutoScalingRepository) DescribeScheduledActions(namespace string) ([]*applicationautoscaling.ScheduledAction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
