package catalogs

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

// NewDeleteCatalogItemsUUIDParams creates a new DeleteCatalogItemsUUIDParams object
// with the default values initialized.
func NewDeleteCatalogItemsUUIDParams() *DeleteCatalogItemsUUIDParams {
	var ()
	return &DeleteCatalogItemsUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCatalogItemsUUIDParamsWithTimeout creates a new DeleteCatalogItemsUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCatalogItemsUUIDParamsWithTimeout(timeout time.Duration) *DeleteCatalogItemsUUIDParams {
	var ()
	return &DeleteCatalogItemsUUIDParams{

		timeout: timeout,
	}
}

// NewDeleteCatalogItemsUUIDParamsWithContext creates a new DeleteCatalogItemsUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCatalogItemsUUIDParamsWithContext(ctx context.Context) *DeleteCatalogItemsUUIDParams {
	var ()
	return &DeleteCatalogItemsUUIDParams{

		Context: ctx,
	}
}

/*DeleteCatalogItemsUUIDParams contains all the parameters to send to the API endpoint
for the delete catalog items UUID operation typically these are written to a http.Request
*/
type DeleteCatalogItemsUUIDParams struct {

	/*UUID
	  The UUID of the entity.

	*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the delete catalog items UUID params
func (o *DeleteCatalogItemsUUIDParams) WithTimeout(timeout time.Duration) *DeleteCatalogItemsUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete catalog items UUID params
func (o *DeleteCatalogItemsUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete catalog items UUID params
func (o *DeleteCatalogItemsUUIDParams) WithContext(ctx context.Context) *DeleteCatalogItemsUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete catalog items UUID params
func (o *DeleteCatalogItemsUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithUUID adds the uuid to the delete catalog items UUID params
func (o *DeleteCatalogItemsUUIDParams) WithUUID(uuid string) *DeleteCatalogItemsUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete catalog items UUID params
func (o *DeleteCatalogItemsUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCatalogItemsUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
