// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ipam

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

// NewIpamFhrpGroupsReadParams creates a new IpamFhrpGroupsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamFhrpGroupsReadParams() *IpamFhrpGroupsReadParams {
	return &IpamFhrpGroupsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamFhrpGroupsReadParamsWithTimeout creates a new IpamFhrpGroupsReadParams object
// with the ability to set a timeout on a request.
func NewIpamFhrpGroupsReadParamsWithTimeout(timeout time.Duration) *IpamFhrpGroupsReadParams {
	return &IpamFhrpGroupsReadParams{
		timeout: timeout,
	}
}

// NewIpamFhrpGroupsReadParamsWithContext creates a new IpamFhrpGroupsReadParams object
// with the ability to set a context for a request.
func NewIpamFhrpGroupsReadParamsWithContext(ctx context.Context) *IpamFhrpGroupsReadParams {
	return &IpamFhrpGroupsReadParams{
		Context: ctx,
	}
}

// NewIpamFhrpGroupsReadParamsWithHTTPClient creates a new IpamFhrpGroupsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamFhrpGroupsReadParamsWithHTTPClient(client *http.Client) *IpamFhrpGroupsReadParams {
	return &IpamFhrpGroupsReadParams{
		HTTPClient: client,
	}
}

/* IpamFhrpGroupsReadParams contains all the parameters to send to the API endpoint
   for the ipam fhrp groups read operation.

   Typically these are written to a http.Request.
*/
type IpamFhrpGroupsReadParams struct {

	/* ID.

	   A unique integer value identifying this FHRP group.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam fhrp groups read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamFhrpGroupsReadParams) WithDefaults() *IpamFhrpGroupsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam fhrp groups read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamFhrpGroupsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam fhrp groups read params
func (o *IpamFhrpGroupsReadParams) WithTimeout(timeout time.Duration) *IpamFhrpGroupsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam fhrp groups read params
func (o *IpamFhrpGroupsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam fhrp groups read params
func (o *IpamFhrpGroupsReadParams) WithContext(ctx context.Context) *IpamFhrpGroupsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam fhrp groups read params
func (o *IpamFhrpGroupsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam fhrp groups read params
func (o *IpamFhrpGroupsReadParams) WithHTTPClient(client *http.Client) *IpamFhrpGroupsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam fhrp groups read params
func (o *IpamFhrpGroupsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam fhrp groups read params
func (o *IpamFhrpGroupsReadParams) WithID(id int64) *IpamFhrpGroupsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam fhrp groups read params
func (o *IpamFhrpGroupsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamFhrpGroupsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}