package packet_processor_chain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeletePacketProcessorChainsUUIDParams creates a new DeletePacketProcessorChainsUUIDParams object
// with the default values initialized.
func NewDeletePacketProcessorChainsUUIDParams() *DeletePacketProcessorChainsUUIDParams {
	var ()
	return &DeletePacketProcessorChainsUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePacketProcessorChainsUUIDParamsWithTimeout creates a new DeletePacketProcessorChainsUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePacketProcessorChainsUUIDParamsWithTimeout(timeout time.Duration) *DeletePacketProcessorChainsUUIDParams {
	var ()
	return &DeletePacketProcessorChainsUUIDParams{

		timeout: timeout,
	}
}

// NewDeletePacketProcessorChainsUUIDParamsWithContext creates a new DeletePacketProcessorChainsUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePacketProcessorChainsUUIDParamsWithContext(ctx context.Context) *DeletePacketProcessorChainsUUIDParams {
	var ()
	return &DeletePacketProcessorChainsUUIDParams{

		Context: ctx,
	}
}

/*DeletePacketProcessorChainsUUIDParams contains all the parameters to send to the API endpoint
for the delete packet processor chains UUID operation typically these are written to a http.Request
*/
type DeletePacketProcessorChainsUUIDParams struct {

	/*UUID*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the delete packet processor chains UUID params
func (o *DeletePacketProcessorChainsUUIDParams) WithTimeout(timeout time.Duration) *DeletePacketProcessorChainsUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete packet processor chains UUID params
func (o *DeletePacketProcessorChainsUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete packet processor chains UUID params
func (o *DeletePacketProcessorChainsUUIDParams) WithContext(ctx context.Context) *DeletePacketProcessorChainsUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete packet processor chains UUID params
func (o *DeletePacketProcessorChainsUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithUUID adds the uuid to the delete packet processor chains UUID params
func (o *DeletePacketProcessorChainsUUIDParams) WithUUID(uuid string) *DeletePacketProcessorChainsUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete packet processor chains UUID params
func (o *DeletePacketProcessorChainsUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePacketProcessorChainsUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
