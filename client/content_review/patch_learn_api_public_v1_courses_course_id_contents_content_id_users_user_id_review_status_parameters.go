// Code generated by go-swagger; DO NOT EDIT.

package content_review

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

// NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams creates a new PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams object
// with the default values initialized.
func NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams() *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams {
	var ()
	return &PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParamsWithTimeout creates a new PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParamsWithTimeout(timeout time.Duration) *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams {
	var ()
	return &PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams{

		timeout: timeout,
	}
}

// NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParamsWithContext creates a new PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParamsWithContext(ctx context.Context) *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams {
	var ()
	return &PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams{

		Context: ctx,
	}
}

// NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParamsWithHTTPClient creates a new PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParamsWithHTTPClient(client *http.Client) *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams {
	var ()
	return &PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams{
		HTTPClient: client,
	}
}

/*PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams contains all the parameters to send to the API endpoint
for the patch learn API public v1 courses course ID contents content ID users user ID review status operation typically these are written to a http.Request
*/
type PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams struct {

	/*ContentID
	 The Content ID.  This may be the primary ID, or any of the following keywords: interactive, indirect, root.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| keyword    | root                                  |


	*/
	ContentID string
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
	/*Input*/
	Input PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusBody
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

// WithTimeout adds the timeout to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) WithTimeout(timeout time.Duration) *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) WithContext(ctx context.Context) *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) WithHTTPClient(client *http.Client) *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentID adds the contentID to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) WithContentID(contentID string) *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams {
	o.SetContentID(contentID)
	return o
}

// SetContentID adds the contentId to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) SetContentID(contentID string) {
	o.ContentID = contentID
}

// WithCourseID adds the courseID to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) WithCourseID(courseID string) *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) WithFields(fields *string) *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithInput adds the input to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) WithInput(input PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusBody) *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) SetInput(input PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusBody) {
	o.Input = input
}

// WithUserID adds the userID to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) WithUserID(userID string) *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the patch learn API public v1 courses course ID contents content ID users user ID review status params
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLearnAPIPublicV1CoursesCourseIDContentsContentIDUsersUserIDReviewStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contentId
	if err := r.SetPathParam("contentId", o.ContentID); err != nil {
		return err
	}

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
