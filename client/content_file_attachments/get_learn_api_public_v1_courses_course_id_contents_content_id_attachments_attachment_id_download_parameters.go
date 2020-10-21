// Code generated by go-swagger; DO NOT EDIT.

package content_file_attachments

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

// NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams creates a new GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams() *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParamsWithTimeout creates a new GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParamsWithContext creates a new GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParamsWithHTTPClient creates a new GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams contains all the parameters to send to the API endpoint
for the get learn API public v1 courses course ID contents content ID attachments attachment ID download operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams struct {

	/*AttachmentID*/
	AttachmentID string
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 courses course ID contents content ID attachments attachment ID download params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 courses course ID contents content ID attachments attachment ID download params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 courses course ID contents content ID attachments attachment ID download params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 courses course ID contents content ID attachments attachment ID download params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID contents content ID attachments attachment ID download params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID contents content ID attachments attachment ID download params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttachmentID adds the attachmentID to the get learn API public v1 courses course ID contents content ID attachments attachment ID download params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams) WithAttachmentID(attachmentID string) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams {
	o.SetAttachmentID(attachmentID)
	return o
}

// SetAttachmentID adds the attachmentId to the get learn API public v1 courses course ID contents content ID attachments attachment ID download params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams) SetAttachmentID(attachmentID string) {
	o.AttachmentID = attachmentID
}

// WithContentID adds the contentID to the get learn API public v1 courses course ID contents content ID attachments attachment ID download params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams) WithContentID(contentID string) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams {
	o.SetContentID(contentID)
	return o
}

// SetContentID adds the contentId to the get learn API public v1 courses course ID contents content ID attachments attachment ID download params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams) SetContentID(contentID string) {
	o.ContentID = contentID
}

// WithCourseID adds the courseID to the get learn API public v1 courses course ID contents content ID attachments attachment ID download params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams) WithCourseID(courseID string) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v1 courses course ID contents content ID attachments attachment ID download params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDAttachmentsAttachmentIDDownloadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attachmentId
	if err := r.SetPathParam("attachmentId", o.AttachmentID); err != nil {
		return err
	}

	// path param contentId
	if err := r.SetPathParam("contentId", o.ContentID); err != nil {
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
