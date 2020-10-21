// Code generated by go-swagger; DO NOT EDIT.

package course_categories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams creates a new GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams() *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams {
	var ()
	return &GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParamsWithTimeout creates a new GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams {
	var ()
	return &GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParamsWithContext creates a new GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams {
	var ()
	return &GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParamsWithHTTPClient creates a new GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams {
	var ()
	return &GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams contains all the parameters to send to the API endpoint
for the get learn API public v1 catalog categories category type category ID operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams struct {

	/*CategoryID*/
	CategoryID string
	/*CategoryType*/
	CategoryType string
	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 catalog categories category type category ID params
func (o *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 catalog categories category type category ID params
func (o *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 catalog categories category type category ID params
func (o *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 catalog categories category type category ID params
func (o *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 catalog categories category type category ID params
func (o *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 catalog categories category type category ID params
func (o *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategoryID adds the categoryID to the get learn API public v1 catalog categories category type category ID params
func (o *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams) WithCategoryID(categoryID string) *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the get learn API public v1 catalog categories category type category ID params
func (o *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams) SetCategoryID(categoryID string) {
	o.CategoryID = categoryID
}

// WithCategoryType adds the categoryType to the get learn API public v1 catalog categories category type category ID params
func (o *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams) WithCategoryType(categoryType string) *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams {
	o.SetCategoryType(categoryType)
	return o
}

// SetCategoryType adds the categoryType to the get learn API public v1 catalog categories category type category ID params
func (o *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams) SetCategoryType(categoryType string) {
	o.CategoryType = categoryType
}

// WithFields adds the fields to the get learn API public v1 catalog categories category type category ID params
func (o *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams) WithFields(fields *string) *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 catalog categories category type category ID params
func (o *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param categoryId
	if err := r.SetPathParam("categoryId", o.CategoryID); err != nil {
		return err
	}

	// path param categoryType
	if err := r.SetPathParam("categoryType", o.CategoryType); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}