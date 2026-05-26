package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
	"github.com/snyk/driftctl/enumeration/remote/cache"
)

type AutoScalingRepository interface {
	DescribeLaunchConfigurations() ([]*autoscaling.LaunchConfiguration, error)
}

type autoScalingRepository struct {
	client autoscalingiface.AutoScalingAPI
	cache  cache.Cache
}

func NewAutoScalingRepository(session *session.Session, c cache.Cache) *autoScalingRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *autoScalingRepository) DescribeLaunchConfigurations() ([]*autoscaling.LaunchConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
