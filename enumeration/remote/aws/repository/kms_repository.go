package repository

import (
	"sync"

	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
)

type KMSRepository interface {
	ListAllKeys() ([]*kms.KeyListEntry, error)
	ListAllAliases() ([]*kms.AliasListEntry, error)
}

type kmsRepository struct {
	client          kmsiface.KMSAPI
	cache           cache.Cache
	describeKeyLock *sync.Mutex
}

func NewKMSRepository(session *session.Session, c cache.Cache) *kmsRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *kmsRepository) ListAllKeys() ([]*kms.KeyListEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *kmsRepository) ListAllAliases() ([]*kms.AliasListEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *kmsRepository) describeKey(keyId *string) (*kms.DescribeKeyOutput, error) {
	_ = "STUB: not implemented"
	return nil,

		// Since this method can be call in parallel, we should lock and unlock if we want to be sure to hit the cache
		nil
}

func (r *kmsRepository) filterKeys(keys []*kms.KeyListEntry) ([]*kms.KeyListEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *kmsRepository) filterAliases(aliases []*kms.AliasListEntry) ([]*kms.AliasListEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
