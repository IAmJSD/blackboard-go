// Code generated by go-swagger; DO NOT EDIT.

package data_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new data sources API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for data sources API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteLearnAPIPublicV1DataSourcesDataSourceID(params *DeleteLearnAPIPublicV1DataSourcesDataSourceIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLearnAPIPublicV1DataSourcesDataSourceIDNoContent, error)

	GetLearnAPIPublicV1DataSources(params *GetLearnAPIPublicV1DataSourcesParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1DataSourcesOK, error)

	GetLearnAPIPublicV1DataSourcesDataSourceID(params *GetLearnAPIPublicV1DataSourcesDataSourceIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1DataSourcesDataSourceIDOK, error)

	PatchLearnAPIPublicV1DataSourcesDataSourceID(params *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchLearnAPIPublicV1DataSourcesDataSourceIDOK, error)

	PostLearnAPIPublicV1DataSources(params *PostLearnAPIPublicV1DataSourcesParams, authInfo runtime.ClientAuthInfoWriter) (*PostLearnAPIPublicV1DataSourcesCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteLearnAPIPublicV1DataSourcesDataSourceID deletes data source

  Deletes a data source.

The 'system.datasource.manager.VIEW' entitlement is needed.

**Since**: 2015.11.0
*/
func (a *Client) DeleteLearnAPIPublicV1DataSourcesDataSourceID(params *DeleteLearnAPIPublicV1DataSourcesDataSourceIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLearnAPIPublicV1DataSourcesDataSourceIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLearnAPIPublicV1DataSourcesDataSourceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLearnAPIPublicV1DataSourcesDataSourceID",
		Method:             "DELETE",
		PathPattern:        "/learn/api/public/v1/dataSources/{dataSourceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLearnAPIPublicV1DataSourcesDataSourceIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteLearnAPIPublicV1DataSourcesDataSourceIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteLearnAPIPublicV1DataSourcesDataSourceID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLearnAPIPublicV1DataSources gets data sources

  Returns a list of data sources.

The 'system.datasource.manager.VIEW' entitlement is needed.

**Since**: 3000.1.0
*/
func (a *Client) GetLearnAPIPublicV1DataSources(params *GetLearnAPIPublicV1DataSourcesParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1DataSourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLearnAPIPublicV1DataSourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLearnAPIPublicV1DataSources",
		Method:             "GET",
		PathPattern:        "/learn/api/public/v1/dataSources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearnAPIPublicV1DataSourcesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLearnAPIPublicV1DataSourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLearnAPIPublicV1DataSources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLearnAPIPublicV1DataSourcesDataSourceID gets data source

  Loads a data source.

The 'system.datasource.manager.VIEW' entitlement is needed.

**Since**: 2015.11.0
*/
func (a *Client) GetLearnAPIPublicV1DataSourcesDataSourceID(params *GetLearnAPIPublicV1DataSourcesDataSourceIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1DataSourcesDataSourceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLearnAPIPublicV1DataSourcesDataSourceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLearnAPIPublicV1DataSourcesDataSourceID",
		Method:             "GET",
		PathPattern:        "/learn/api/public/v1/dataSources/{dataSourceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearnAPIPublicV1DataSourcesDataSourceIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLearnAPIPublicV1DataSourcesDataSourceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLearnAPIPublicV1DataSourcesDataSourceID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchLearnAPIPublicV1DataSourcesDataSourceID updates data source

  Updates a data source.

The 'system.datasource.manager.VIEW' entitlement is needed.

**Since**: 2015.11.0
*/
func (a *Client) PatchLearnAPIPublicV1DataSourcesDataSourceID(params *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchLearnAPIPublicV1DataSourcesDataSourceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchLearnAPIPublicV1DataSourcesDataSourceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchLearnAPIPublicV1DataSourcesDataSourceID",
		Method:             "PATCH",
		PathPattern:        "/learn/api/public/v1/dataSources/{dataSourceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchLearnAPIPublicV1DataSourcesDataSourceIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchLearnAPIPublicV1DataSourcesDataSourceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchLearnAPIPublicV1DataSourcesDataSourceID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostLearnAPIPublicV1DataSources creates data source

  Creates a data source.

The 'system.datasource.manager.VIEW' entitlement is needed.

**Since**: 2015.11.0
*/
func (a *Client) PostLearnAPIPublicV1DataSources(params *PostLearnAPIPublicV1DataSourcesParams, authInfo runtime.ClientAuthInfoWriter) (*PostLearnAPIPublicV1DataSourcesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLearnAPIPublicV1DataSourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLearnAPIPublicV1DataSources",
		Method:             "POST",
		PathPattern:        "/learn/api/public/v1/dataSources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLearnAPIPublicV1DataSourcesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLearnAPIPublicV1DataSourcesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostLearnAPIPublicV1DataSources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}