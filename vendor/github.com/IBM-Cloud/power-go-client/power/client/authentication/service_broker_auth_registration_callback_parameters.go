// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// NewServiceBrokerAuthRegistrationCallbackParams creates a new ServiceBrokerAuthRegistrationCallbackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServiceBrokerAuthRegistrationCallbackParams() *ServiceBrokerAuthRegistrationCallbackParams {
	return &ServiceBrokerAuthRegistrationCallbackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServiceBrokerAuthRegistrationCallbackParamsWithTimeout creates a new ServiceBrokerAuthRegistrationCallbackParams object
// with the ability to set a timeout on a request.
func NewServiceBrokerAuthRegistrationCallbackParamsWithTimeout(timeout time.Duration) *ServiceBrokerAuthRegistrationCallbackParams {
	return &ServiceBrokerAuthRegistrationCallbackParams{
		timeout: timeout,
	}
}

// NewServiceBrokerAuthRegistrationCallbackParamsWithContext creates a new ServiceBrokerAuthRegistrationCallbackParams object
// with the ability to set a context for a request.
func NewServiceBrokerAuthRegistrationCallbackParamsWithContext(ctx context.Context) *ServiceBrokerAuthRegistrationCallbackParams {
	return &ServiceBrokerAuthRegistrationCallbackParams{
		Context: ctx,
	}
}

// NewServiceBrokerAuthRegistrationCallbackParamsWithHTTPClient creates a new ServiceBrokerAuthRegistrationCallbackParams object
// with the ability to set a custom HTTPClient for a request.
func NewServiceBrokerAuthRegistrationCallbackParamsWithHTTPClient(client *http.Client) *ServiceBrokerAuthRegistrationCallbackParams {
	return &ServiceBrokerAuthRegistrationCallbackParams{
		HTTPClient: client,
	}
}

/*
ServiceBrokerAuthRegistrationCallbackParams contains all the parameters to send to the API endpoint

	for the service broker auth registration callback operation.

	Typically these are written to a http.Request.
*/
type ServiceBrokerAuthRegistrationCallbackParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service broker auth registration callback params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceBrokerAuthRegistrationCallbackParams) WithDefaults() *ServiceBrokerAuthRegistrationCallbackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service broker auth registration callback params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceBrokerAuthRegistrationCallbackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service broker auth registration callback params
func (o *ServiceBrokerAuthRegistrationCallbackParams) WithTimeout(timeout time.Duration) *ServiceBrokerAuthRegistrationCallbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service broker auth registration callback params
func (o *ServiceBrokerAuthRegistrationCallbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service broker auth registration callback params
func (o *ServiceBrokerAuthRegistrationCallbackParams) WithContext(ctx context.Context) *ServiceBrokerAuthRegistrationCallbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service broker auth registration callback params
func (o *ServiceBrokerAuthRegistrationCallbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service broker auth registration callback params
func (o *ServiceBrokerAuthRegistrationCallbackParams) WithHTTPClient(client *http.Client) *ServiceBrokerAuthRegistrationCallbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service broker auth registration callback params
func (o *ServiceBrokerAuthRegistrationCallbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceBrokerAuthRegistrationCallbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
