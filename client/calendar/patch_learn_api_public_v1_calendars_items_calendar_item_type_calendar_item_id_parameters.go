// Code generated by go-swagger; DO NOT EDIT.

package calendar

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
	"github.com/go-openapi/swag"
)

// NewPatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams creates a new PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams object
// with the default values initialized.
func NewPatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams() *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams {
	var ()
	return &PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParamsWithTimeout creates a new PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParamsWithTimeout(timeout time.Duration) *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams {
	var ()
	return &PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams{

		timeout: timeout,
	}
}

// NewPatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParamsWithContext creates a new PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParamsWithContext(ctx context.Context) *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams {
	var ()
	return &PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams{

		Context: ctx,
	}
}

// NewPatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParamsWithHTTPClient creates a new PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParamsWithHTTPClient(client *http.Client) *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams {
	var ()
	return &PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams{
		HTTPClient: client,
	}
}

/*PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams contains all the parameters to send to the API endpoint
for the patch learn API public v1 calendars items calendar item type calendar item ID operation typically these are written to a http.Request
*/
type PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams struct {

	/*CalendarItemID*/
	CalendarItemID string
	/*CalendarItemType*/
	CalendarItemType string
	/*Input*/
	Input PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDBody
	/*UpdateSeries
	  update the series calendar items or just one calendar item. true - update the        entire series, false - update a single calendar item. Defaults to false. When updating an entire series the full recurrence        object must be populated just as if creating a new calendar series. If updating a single calendar entry the recurrence        must not be specified.

	*/
	UpdateSeries *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch learn API public v1 calendars items calendar item type calendar item ID params
func (o *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams) WithTimeout(timeout time.Duration) *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch learn API public v1 calendars items calendar item type calendar item ID params
func (o *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch learn API public v1 calendars items calendar item type calendar item ID params
func (o *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams) WithContext(ctx context.Context) *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch learn API public v1 calendars items calendar item type calendar item ID params
func (o *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch learn API public v1 calendars items calendar item type calendar item ID params
func (o *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams) WithHTTPClient(client *http.Client) *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch learn API public v1 calendars items calendar item type calendar item ID params
func (o *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCalendarItemID adds the calendarItemID to the patch learn API public v1 calendars items calendar item type calendar item ID params
func (o *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams) WithCalendarItemID(calendarItemID string) *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams {
	o.SetCalendarItemID(calendarItemID)
	return o
}

// SetCalendarItemID adds the calendarItemId to the patch learn API public v1 calendars items calendar item type calendar item ID params
func (o *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams) SetCalendarItemID(calendarItemID string) {
	o.CalendarItemID = calendarItemID
}

// WithCalendarItemType adds the calendarItemType to the patch learn API public v1 calendars items calendar item type calendar item ID params
func (o *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams) WithCalendarItemType(calendarItemType string) *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams {
	o.SetCalendarItemType(calendarItemType)
	return o
}

// SetCalendarItemType adds the calendarItemType to the patch learn API public v1 calendars items calendar item type calendar item ID params
func (o *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams) SetCalendarItemType(calendarItemType string) {
	o.CalendarItemType = calendarItemType
}

// WithInput adds the input to the patch learn API public v1 calendars items calendar item type calendar item ID params
func (o *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams) WithInput(input PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDBody) *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the patch learn API public v1 calendars items calendar item type calendar item ID params
func (o *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams) SetInput(input PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDBody) {
	o.Input = input
}

// WithUpdateSeries adds the updateSeries to the patch learn API public v1 calendars items calendar item type calendar item ID params
func (o *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams) WithUpdateSeries(updateSeries *bool) *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams {
	o.SetUpdateSeries(updateSeries)
	return o
}

// SetUpdateSeries adds the updateSeries to the patch learn API public v1 calendars items calendar item type calendar item ID params
func (o *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams) SetUpdateSeries(updateSeries *bool) {
	o.UpdateSeries = updateSeries
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLearnAPIPublicV1CalendarsItemsCalendarItemTypeCalendarItemIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param calendarItemId
	if err := r.SetPathParam("calendarItemId", o.CalendarItemID); err != nil {
		return err
	}

	// path param calendarItemType
	if err := r.SetPathParam("calendarItemType", o.CalendarItemType); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Input); err != nil {
		return err
	}

	if o.UpdateSeries != nil {

		// query param updateSeries
		var qrUpdateSeries bool
		if o.UpdateSeries != nil {
			qrUpdateSeries = *o.UpdateSeries
		}
		qUpdateSeries := swag.FormatBool(qrUpdateSeries)
		if qUpdateSeries != "" {
			if err := r.SetQueryParam("updateSeries", qUpdateSeries); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}