package repository

import (
	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
)

type IAMRepository interface {
	ListAllAccessKeys([]*iam.User) ([]*iam.AccessKeyMetadata, error)
	ListAllUsers() ([]*iam.User, error)
	ListAllPolicies() ([]*iam.Policy, error)
	ListAllRoles() ([]*iam.Role, error)
	ListAllRolePolicyAttachments([]*iam.Role) ([]*AttachedRolePolicy, error)
	ListAllRolePolicies([]*iam.Role) ([]RolePolicy, error)
	ListAllUserPolicyAttachments([]*iam.User) ([]*AttachedUserPolicy, error)
	ListAllUserPolicies([]*iam.User) ([]string, error)
	ListAllGroups() ([]*iam.Group, error)
	ListAllGroupPolicies([]*iam.Group) ([]string, error)
	ListAllGroupPolicyAttachments([]*iam.Group) ([]*AttachedGroupPolicy, error)
}

type iamRepository struct {
	client iamiface.IAMAPI
	cache  cache.Cache
}

func NewIAMRepository(session *session.Session, c cache.Cache) *iamRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *iamRepository) ListAllAccessKeys(users []*iam.User) ([]*iam.AccessKeyMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *iamRepository) ListAllUsers() ([]*iam.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *iamRepository) ListAllPolicies() ([]*iam.Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *iamRepository) ListAllRoles() ([]*iam.Role, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *iamRepository) ListAllRolePolicyAttachments(roles []*iam.Role) ([]*AttachedRolePolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *iamRepository) ListAllRolePolicies(roles []*iam.Role) ([]RolePolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *iamRepository) ListAllUserPolicyAttachments(users []*iam.User) ([]*AttachedUserPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *iamRepository) ListAllUserPolicies(users []*iam.User) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *iamRepository) ListAllGroups() ([]*iam.Group, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *iamRepository) ListAllGroupPolicies(groups []*iam.Group) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *iamRepository) ListAllGroupPolicyAttachments(groups []*iam.Group) ([]*AttachedGroupPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type AttachedUserPolicy struct {
	iam.AttachedPolicy
	UserName string
}

type AttachedRolePolicy struct {
	iam.AttachedPolicy
	RoleName string
}

type AttachedGroupPolicy struct {
	iam.AttachedPolicy
	GroupName string
}

type RolePolicy struct {
	Policy   string
	RoleName string
}
