// Code generated by go-swagger; DO NOT EDIT.

package course_grades

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

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams creates a new GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams() *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParamsWithTimeout creates a new GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParamsWithContext creates a new GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParamsWithHTTPClient creates a new GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams contains all the parameters to send to the API endpoint
for the get learn API public v1 courses course ID gradebook schemas schema ID operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams struct {

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
	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*SchemaID
	 The grade schema ID.  This may be the primary ID, or any of the secondary IDs prefixed with the ID type.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| externalId | externalId:schema1                    |


	*/
	SchemaID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 courses course ID gradebook schemas schema ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 courses course ID gradebook schemas schema ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 courses course ID gradebook schemas schema ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 courses course ID gradebook schemas schema ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID gradebook schemas schema ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID gradebook schemas schema ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCourseID adds the courseID to the get learn API public v1 courses course ID gradebook schemas schema ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams) WithCourseID(courseID string) *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v1 courses course ID gradebook schemas schema ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the get learn API public v1 courses course ID gradebook schemas schema ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams) WithFields(fields *string) *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 courses course ID gradebook schemas schema ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithSchemaID adds the schemaID to the get learn API public v1 courses course ID gradebook schemas schema ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams) WithSchemaID(schemaID string) *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams {
	o.SetSchemaID(schemaID)
	return o
}

// SetSchemaID adds the schemaId to the get learn API public v1 courses course ID gradebook schemas schema ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams) SetSchemaID(schemaID string) {
	o.SchemaID = schemaID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param courseId
	if err := r.SetPathParam("courseId", o.CourseID); err != nil {
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

	// path param schemaId
	if err := r.SetPathParam("schemaId", o.SchemaID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
