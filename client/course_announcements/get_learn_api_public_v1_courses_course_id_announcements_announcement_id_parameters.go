// Code generated by go-swagger; DO NOT EDIT.

package course_announcements

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

// NewGetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams creates a new GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams() *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParamsWithTimeout creates a new GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParamsWithContext creates a new GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParamsWithHTTPClient creates a new GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams contains all the parameters to send to the API endpoint
for the get learn API public v1 courses course ID announcements announcement ID operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams struct {

	/*AnnouncementID*/
	AnnouncementID string
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 courses course ID announcements announcement ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 courses course ID announcements announcement ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 courses course ID announcements announcement ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 courses course ID announcements announcement ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID announcements announcement ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID announcements announcement ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAnnouncementID adds the announcementID to the get learn API public v1 courses course ID announcements announcement ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams) WithAnnouncementID(announcementID string) *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams {
	o.SetAnnouncementID(announcementID)
	return o
}

// SetAnnouncementID adds the announcementId to the get learn API public v1 courses course ID announcements announcement ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams) SetAnnouncementID(announcementID string) {
	o.AnnouncementID = announcementID
}

// WithCourseID adds the courseID to the get learn API public v1 courses course ID announcements announcement ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams) WithCourseID(courseID string) *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v1 courses course ID announcements announcement ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the get learn API public v1 courses course ID announcements announcement ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams) WithFields(fields *string) *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 courses course ID announcements announcement ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param announcementId
	if err := r.SetPathParam("announcementId", o.AnnouncementID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}