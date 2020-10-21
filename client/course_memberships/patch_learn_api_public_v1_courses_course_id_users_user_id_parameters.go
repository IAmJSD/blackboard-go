// Code generated by go-swagger; DO NOT EDIT.

package course_memberships

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

// NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams creates a new PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams object
// with the default values initialized.
func NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams() *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams {
	var ()
	return &PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParamsWithTimeout creates a new PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParamsWithTimeout(timeout time.Duration) *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams {
	var ()
	return &PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams{

		timeout: timeout,
	}
}

// NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParamsWithContext creates a new PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParamsWithContext(ctx context.Context) *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams {
	var ()
	return &PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams{

		Context: ctx,
	}
}

// NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParamsWithHTTPClient creates a new PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParamsWithHTTPClient(client *http.Client) *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams {
	var ()
	return &PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams{
		HTTPClient: client,
	}
}

/*PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams contains all the parameters to send to the API endpoint
for the patch learn API public v1 courses course ID users user ID operation typically these are written to a http.Request
*/
type PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams struct {

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
	/*Input
	  JSON input object.

	*/
	Input PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBody
	/*UserID
	 The user ID.  This may be the primary ID, or any of the secondary IDs prefixed with the ID type.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| externalId | externalId:jsmith                     |
	| userName   | userName:jsmith                       |
	| uuid       | uuid:915c7567d76d444abf1eed56aad3beb5 |


	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch learn API public v1 courses course ID users user ID params
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams) WithTimeout(timeout time.Duration) *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch learn API public v1 courses course ID users user ID params
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch learn API public v1 courses course ID users user ID params
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams) WithContext(ctx context.Context) *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch learn API public v1 courses course ID users user ID params
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch learn API public v1 courses course ID users user ID params
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams) WithHTTPClient(client *http.Client) *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch learn API public v1 courses course ID users user ID params
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCourseID adds the courseID to the patch learn API public v1 courses course ID users user ID params
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams) WithCourseID(courseID string) *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the patch learn API public v1 courses course ID users user ID params
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the patch learn API public v1 courses course ID users user ID params
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams) WithFields(fields *string) *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the patch learn API public v1 courses course ID users user ID params
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithInput adds the input to the patch learn API public v1 courses course ID users user ID params
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams) WithInput(input PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBody) *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the patch learn API public v1 courses course ID users user ID params
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams) SetInput(input PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBody) {
	o.Input = input
}

// WithUserID adds the userID to the patch learn API public v1 courses course ID users user ID params
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams) WithUserID(userID string) *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the patch learn API public v1 courses course ID users user ID params
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.Input); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
