// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/e2b-dev/infra/packages/env-instance-task-driver/internal/client/models"
)

// NewPutBalloonParams creates a new PutBalloonParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutBalloonParams() *PutBalloonParams {
	return &PutBalloonParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutBalloonParamsWithTimeout creates a new PutBalloonParams object
// with the ability to set a timeout on a request.
func NewPutBalloonParamsWithTimeout(timeout time.Duration) *PutBalloonParams {
	return &PutBalloonParams{
		timeout: timeout,
	}
}

// NewPutBalloonParamsWithContext creates a new PutBalloonParams object
// with the ability to set a context for a request.
func NewPutBalloonParamsWithContext(ctx context.Context) *PutBalloonParams {
	return &PutBalloonParams{
		Context: ctx,
	}
}

// NewPutBalloonParamsWithHTTPClient creates a new PutBalloonParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutBalloonParamsWithHTTPClient(client *http.Client) *PutBalloonParams {
	return &PutBalloonParams{
		HTTPClient: client,
	}
}

/*
PutBalloonParams contains all the parameters to send to the API endpoint

	for the put balloon operation.

	Typically these are written to a http.Request.
*/
type PutBalloonParams struct {

	/* Body.

	   Balloon properties
	*/
	Body *models.Balloon

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put balloon params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutBalloonParams) WithDefaults() *PutBalloonParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put balloon params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutBalloonParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put balloon params
func (o *PutBalloonParams) WithTimeout(timeout time.Duration) *PutBalloonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put balloon params
func (o *PutBalloonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put balloon params
func (o *PutBalloonParams) WithContext(ctx context.Context) *PutBalloonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put balloon params
func (o *PutBalloonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put balloon params
func (o *PutBalloonParams) WithHTTPClient(client *http.Client) *PutBalloonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put balloon params
func (o *PutBalloonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put balloon params
func (o *PutBalloonParams) WithBody(body *models.Balloon) *PutBalloonParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put balloon params
func (o *PutBalloonParams) SetBody(body *models.Balloon) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutBalloonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
