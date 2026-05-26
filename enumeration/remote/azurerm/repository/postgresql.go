package repository

import (
	"context"

	"github.com/snyk/driftctl/enumeration/remote/azurerm/common"
	"github.com/snyk/driftctl/enumeration/remote/cache"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

type PostgresqlRespository interface {
	ListAllServers() ([]*armpostgresql.Server, error)
	ListAllDatabasesByServer(server *armpostgresql.Server) ([]*armpostgresql.Database, error)
}

type postgresqlServersClientImpl struct {
	client *armpostgresql.ServersClient
}

type postgresqlServersClient interface {
	List(context.Context, *armpostgresql.ServersListOptions) (armpostgresql.ServersListResponse, error)
}

func (c postgresqlServersClientImpl) List(ctx context.Context, options *armpostgresql.ServersListOptions) (armpostgresql.ServersListResponse, error) {
	_ = "STUB: not implemented"
	return *new(armpostgresql.ServersListResponse), nil
}

type postgresqlDatabaseClientImpl struct {
	client *armpostgresql.DatabasesClient
}

type postgresqlDatabaseClient interface {
	ListByServer(context.Context, string, string, *armpostgresql.DatabasesListByServerOptions) (armpostgresql.DatabasesListByServerResponse, error)
}

func (c postgresqlDatabaseClientImpl) ListByServer(ctx context.Context, resGroup string, serverName string, options *armpostgresql.DatabasesListByServerOptions) (armpostgresql.DatabasesListByServerResponse, error) {
	_ = "STUB: not implemented"
	return *new(armpostgresql.DatabasesListByServerResponse), nil
}

type postgresqlRepository struct {
	serversClient  postgresqlServersClient
	databaseClient postgresqlDatabaseClient
	cache          cache.Cache
}

func NewPostgresqlRepository(cred azcore.TokenCredential, options *arm.ClientOptions, config common.AzureProviderConfig, cache cache.Cache) *postgresqlRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *postgresqlRepository) ListAllServers() ([]*armpostgresql.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *postgresqlRepository) ListAllDatabasesByServer(server *armpostgresql.Server) ([]*armpostgresql.Database, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
