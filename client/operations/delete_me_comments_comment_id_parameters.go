// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteMeCommentsCommentIDParams creates a new DeleteMeCommentsCommentIDParams object
// with the default values initialized.
func NewDeleteMeCommentsCommentIDParams() *DeleteMeCommentsCommentIDParams {
	var ()
	return &DeleteMeCommentsCommentIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMeCommentsCommentIDParamsWithTimeout creates a new DeleteMeCommentsCommentIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMeCommentsCommentIDParamsWithTimeout(timeout time.Duration) *DeleteMeCommentsCommentIDParams {
	var ()
	return &DeleteMeCommentsCommentIDParams{

		timeout: timeout,
	}
}

// NewDeleteMeCommentsCommentIDParamsWithContext creates a new DeleteMeCommentsCommentIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMeCommentsCommentIDParamsWithContext(ctx context.Context) *DeleteMeCommentsCommentIDParams {
	var ()
	return &DeleteMeCommentsCommentIDParams{

		Context: ctx,
	}
}

// NewDeleteMeCommentsCommentIDParamsWithHTTPClient creates a new DeleteMeCommentsCommentIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMeCommentsCommentIDParamsWithHTTPClient(client *http.Client) *DeleteMeCommentsCommentIDParams {
	var ()
	return &DeleteMeCommentsCommentIDParams{
		HTTPClient: client,
	}
}

/*DeleteMeCommentsCommentIDParams contains all the parameters to send to the API endpoint
for the delete me comments comment ID operation typically these are written to a http.Request
*/
type DeleteMeCommentsCommentIDParams struct {

	/*CommentID
	  対象コメントを指定するために使用

	*/
	CommentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete me comments comment ID params
func (o *DeleteMeCommentsCommentIDParams) WithTimeout(timeout time.Duration) *DeleteMeCommentsCommentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete me comments comment ID params
func (o *DeleteMeCommentsCommentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete me comments comment ID params
func (o *DeleteMeCommentsCommentIDParams) WithContext(ctx context.Context) *DeleteMeCommentsCommentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete me comments comment ID params
func (o *DeleteMeCommentsCommentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete me comments comment ID params
func (o *DeleteMeCommentsCommentIDParams) WithHTTPClient(client *http.Client) *DeleteMeCommentsCommentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete me comments comment ID params
func (o *DeleteMeCommentsCommentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommentID adds the commentID to the delete me comments comment ID params
func (o *DeleteMeCommentsCommentIDParams) WithCommentID(commentID string) *DeleteMeCommentsCommentIDParams {
	o.SetCommentID(commentID)
	return o
}

// SetCommentID adds the commentId to the delete me comments comment ID params
func (o *DeleteMeCommentsCommentIDParams) SetCommentID(commentID string) {
	o.CommentID = commentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMeCommentsCommentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param comment_id
	if err := r.SetPathParam("comment_id", o.CommentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
