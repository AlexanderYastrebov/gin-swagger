// Code generated by gin-swagger; DO NOT EDIT.

package node_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
	"github.com/mikkeloscar/gin-swagger/api"
	"github.com/mikkeloscar/gin-swagger/tracing"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteNodePoolEndpoint executes the core logic of the related
// route endpoint.
func DeleteNodePoolEndpoint(handler func(ctx *gin.Context, params *DeleteNodePoolParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		span := opentracing.SpanFromContext(tracing.Context(ctx))

		// attach tags to opentracing span
		if span != nil {
			ext.HTTPMethod.Set(span, ctx.Request.Method)
			ext.HTTPUrl.Set(span, ctx.Request.URL.String())
		}

		// generate params from request
		params := NewDeleteNodePoolParams()
		err := params.readRequest(ctx)
		if err != nil {
			errObj := err.(*errors.CompositeError)
			problem := api.Problem{
				Title:  "Unprocessable Entity.",
				Status: int(errObj.Code()),
				Detail: errObj.Error(),
			}

			// attach tags to opentracing span
			if span != nil {
				ext.HTTPStatusCode.Set(span, uint16(problem.Status))
			}

			ctx.Writer.Header().Set("Content-Type", "application/problem+json")
			ctx.JSON(problem.Status, problem)
			return
		}

		resp := handler(ctx, params)

		// attach tags to opentracing span
		if span != nil {
			ext.HTTPStatusCode.Set(span, uint16(resp.Code))
		}

		switch resp.Code {
		case http.StatusNoContent:
			ctx.AbortWithStatus(resp.Code)
		default:
			ctx.JSON(resp.Code, resp.Body)
		}
	}
}

// NewDeleteNodePoolParams creates a new DeleteNodePoolParams object
// with the default values initialized.
func NewDeleteNodePoolParams() *DeleteNodePoolParams {
	var ()
	return &DeleteNodePoolParams{}
}

// DeleteNodePoolParams contains all the bound params for the delete node pool operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteNodePool
type DeleteNodePoolParams struct {

	/*ID of the cluster.
	  Required: true
	  Pattern: ^[a-z][a-z0-9-:]*[a-z0-9]$
	  In: path
	*/
	ClusterID string
	/*Name of the node pool.
	  Required: true
	  Pattern: ^[a-z][a-z0-9-]*[a-z0-9]$
	  In: path
	*/
	NodePoolName string
}

// readRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteNodePoolParams) readRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	rClusterID := []string{ctx.Param("cluster_id")}
	if err := o.bindClusterID(rClusterID, true, formats); err != nil {
		res = append(res, err)
	}

	rNodePoolName := []string{ctx.Param("node_pool_name")}
	if err := o.bindNodePoolName(rNodePoolName, true, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteNodePoolParams) bindClusterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ClusterID = raw

	if err := o.validateClusterID(formats); err != nil {
		return err
	}

	return nil
}

func (o *DeleteNodePoolParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Pattern("cluster_id", "path", o.ClusterID, `^[a-z][a-z0-9-:]*[a-z0-9]$`); err != nil {
		return err
	}

	return nil
}

func (o *DeleteNodePoolParams) bindNodePoolName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.NodePoolName = raw

	if err := o.validateNodePoolName(formats); err != nil {
		return err
	}

	return nil
}

func (o *DeleteNodePoolParams) validateNodePoolName(formats strfmt.Registry) error {

	if err := validate.Pattern("node_pool_name", "path", o.NodePoolName, `^[a-z][a-z0-9-]*[a-z0-9]$`); err != nil {
		return err
	}

	return nil
}

// vim: ft=go
