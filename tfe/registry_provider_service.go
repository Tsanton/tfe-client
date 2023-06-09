package tfe

import (
	"context"
	"fmt"
	"net/http"

	mreq "github.com/tsanton/tfe-client/tfe/models/request"
	mresp "github.com/tsanton/tfe-client/tfe/models/response"
	u "github.com/tsanton/tfe-client/tfe/utilities"
)

type RegistryProviderService struct {
	cli    *TerraformEnterpriseClient
	logger u.ILogger
}

func newRegistryProviderService(cli *TerraformEnterpriseClient, logger u.ILogger) *RegistryProviderService {
	return &RegistryProviderService{
		cli:    cli,
		logger: logger,
	}
}

func (s *RegistryProviderService) Create(ctx context.Context, organization string, prov *mreq.Provider) (mresp.Provider, error) {
	path := fmt.Sprintf("/api/v2/organizations/%s/registry-providers", organization)
	resp, err := MakeRequest[*mreq.Provider, mresp.Provider](ctx, s.cli, http.MethodPost, 201, path, prov)
	if err != nil {
		return mresp.Provider{}, err
	}
	return *resp, nil
}

func (s *RegistryProviderService) Read(ctx context.Context, organization, registryName, namespace, providerName string) (mresp.Provider, error) {
	path := fmt.Sprintf("/api/v2/organizations/%s/registry-providers/%s/%s/%s", organization, registryName, namespace, providerName)
	resp, err := MakeRequest[interface{}, mresp.Provider](ctx, s.cli, http.MethodGet, 200, path, nil)
	if err != nil {
		return mresp.Provider{}, err
	}
	return *resp, nil
}

func (s *RegistryProviderService) Update(ctx context.Context) error {
	panic("not implemented")
}

func (s *RegistryProviderService) Delete(ctx context.Context, organization, registryName, namespace, providerName string) error {
	path := fmt.Sprintf("/api/v2/organizations/%s/registry-providers/%s/%s/%s", organization, registryName, namespace, providerName)
	_, err := MakeRequest[interface{}, interface{}](ctx, s.cli, http.MethodDelete, 204, path, nil)
	if err != nil {
		return err
	}
	return nil
}
