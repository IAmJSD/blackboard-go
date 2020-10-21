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

// NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams creates a new DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams object
// with the default values initialized.
func NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams() *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams {
	var ()
	return &DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParamsWithTimeout creates a new DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParamsWithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams {
	var ()
	return &DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams{

		timeout: timeout,
	}
}

// NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParamsWithContext creates a new DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParamsWithContext(ctx context.Context) *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams {
	var ()
	return &DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams{

		Context: ctx,
	}
}

// NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParamsWithHTTPClient creates a new DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParamsWithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams {
	var ()
	return &DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams{
		HTTPClient: client,
	}
}

/*DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams contains all the parameters to send to the API endpoint
for the delete learn API public v1 catalog categories category type category ID courses course ID operation typically these are written to a http.Request
*/
type DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams struct {

	/*CategoryID*/
	CategoryID string
	/*CategoryType*/
	CategoryType string
	/*CourseID
	 The course or organization ID.  This may be the primary ID, or any of the secondary IDs prefixed with the ID type.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| externalId | externalId:math101                    |
	| courseId   | courseId:math101                      |
	| uuid       | uuid:915c7567d76d444abf1eed56aad3beb5 |


	*/
	CourseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete learn API public v1 catalog categories category type category ID courses course ID params
func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams) WithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete learn API public v1 catalog categories category type category ID courses course ID params
func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete learn API public v1 catalog categories category type category ID courses course ID params
func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams) WithContext(ctx context.Context) *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete learn API public v1 catalog categories category type category ID courses course ID params
func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete learn API public v1 catalog categories category type category ID courses course ID params
func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams) WithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete learn API public v1 catalog categories category type category ID courses course ID params
func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategoryID adds the categoryID to the delete learn API public v1 catalog categories category type category ID courses course ID params
func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams) WithCategoryID(categoryID string) *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the delete learn API public v1 catalog categories category type category ID courses course ID params
func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams) SetCategoryID(categoryID string) {
	o.CategoryID = categoryID
}

// WithCategoryType adds the categoryType to the delete learn API public v1 catalog categories category type category ID courses course ID params
func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams) WithCategoryType(categoryType string) *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams {
	o.SetCategoryType(categoryType)
	return o
}

// SetCategoryType adds the categoryType to the delete learn API public v1 catalog categories category type category ID courses course ID params
func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams) SetCategoryType(categoryType string) {
	o.CategoryType = categoryType
}

// WithCourseID adds the courseID to the delete learn API public v1 catalog categories category type category ID courses course ID params
func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams) WithCourseID(courseID string) *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the delete learn API public v1 catalog categories category type category ID courses course ID params
func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param courseId
	if err := r.SetPathParam("courseId", o.CourseID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}