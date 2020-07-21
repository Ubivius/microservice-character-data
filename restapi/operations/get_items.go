// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetItemsHandlerFunc turns a function with the right signature into a get items handler
type GetItemsHandlerFunc func(GetItemsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetItemsHandlerFunc) Handle(params GetItemsParams) middleware.Responder {
	return fn(params)
}

// GetItemsHandler interface for that can handle valid get items params
type GetItemsHandler interface {
	Handle(GetItemsParams) middleware.Responder
}

// NewGetItems creates a new http.Handler for the get items operation
func NewGetItems(ctx *middleware.Context, handler GetItemsHandler) *GetItems {
	return &GetItems{Context: ctx, Handler: handler}
}

/*GetItems swagger:route GET /items getItems

GetItems get items API

*/
type GetItems struct {
	Context *middleware.Context
	Handler GetItemsHandler
}

func (o *GetItems) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetItemsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
