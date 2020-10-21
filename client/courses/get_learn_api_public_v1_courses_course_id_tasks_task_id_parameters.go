// Code generated by go-swagger; DO NOT EDIT.

package courses

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

// NewGetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams creates a new GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams() *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParamsWithTimeout creates a new GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParamsWithContext creates a new GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParamsWithHTTPClient creates a new GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams contains all the parameters to send to the API endpoint
for the get learn API public v1 courses course ID tasks task ID operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams struct {

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
	/*TaskID*/
	TaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 courses course ID tasks task ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 courses course ID tasks task ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 courses course ID tasks task ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 courses course ID tasks task ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID tasks task ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID tasks task ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCourseID adds the courseID to the get learn API public v1 courses course ID tasks task ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams) WithCourseID(courseID string) *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v1 courses course ID tasks task ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the get learn API public v1 courses course ID tasks task ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams) WithFields(fields *string) *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 courses course ID tasks task ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithTaskID adds the taskID to the get learn API public v1 courses course ID tasks task ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams) WithTaskID(taskID string) *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the get learn API public v1 courses course ID tasks task ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1CoursesCourseIDTasksTaskIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param taskId
	if err := r.SetPathParam("taskId", o.TaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
