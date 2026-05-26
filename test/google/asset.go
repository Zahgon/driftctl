package google

import (
	"context"

	asset "cloud.google.com/go/asset/apiv1"
	assetpb "cloud.google.com/go/asset/apiv1/assetpb"
)

type FakeAssetServer struct {
	SearchAllResourcesResults []*assetpb.ResourceSearchResult
	ListAssetsResults         []*assetpb.Asset
	err                       error
	assetpb.UnimplementedAssetServiceServer
}

func (s *FakeAssetServer) SearchAllResources(context.Context, *assetpb.SearchAllResourcesRequest) (*assetpb.SearchAllResourcesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *FakeAssetServer) ListAssets(context.Context, *assetpb.ListAssetsRequest) (*assetpb.ListAssetsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFakeAssertServerWithList(listResults []*assetpb.Asset, err error) (*asset.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFakeAssetServer(searchResults []*assetpb.ResourceSearchResult, err error) (*asset.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newAssetClient(fakeServer *FakeAssetServer) (*asset.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create a client.
