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

	models "github.com/alisista/go-alis-client/models"
)

// NewPutMeArticlesArticleIDPublicRepublishParams creates a new PutMeArticlesArticleIDPublicRepublishParams object
// with the default values initialized.
func NewPutMeArticlesArticleIDPublicRepublishParams() *PutMeArticlesArticleIDPublicRepublishParams {
	var ()
	return &PutMeArticlesArticleIDPublicRepublishParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutMeArticlesArticleIDPublicRepublishParamsWithTimeout creates a new PutMeArticlesArticleIDPublicRepublishParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutMeArticlesArticleIDPublicRepublishParamsWithTimeout(timeout time.Duration) *PutMeArticlesArticleIDPublicRepublishParams {
	var ()
	return &PutMeArticlesArticleIDPublicRepublishParams{

		timeout: timeout,
	}
}

// NewPutMeArticlesArticleIDPublicRepublishParamsWithContext creates a new PutMeArticlesArticleIDPublicRepublishParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutMeArticlesArticleIDPublicRepublishParamsWithContext(ctx context.Context) *PutMeArticlesArticleIDPublicRepublishParams {
	var ()
	return &PutMeArticlesArticleIDPublicRepublishParams{

		Context: ctx,
	}
}

// NewPutMeArticlesArticleIDPublicRepublishParamsWithHTTPClient creates a new PutMeArticlesArticleIDPublicRepublishParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutMeArticlesArticleIDPublicRepublishParamsWithHTTPClient(client *http.Client) *PutMeArticlesArticleIDPublicRepublishParams {
	var ()
	return &PutMeArticlesArticleIDPublicRepublishParams{
		HTTPClient: client,
	}
}

/*PutMeArticlesArticleIDPublicRepublishParams contains all the parameters to send to the API endpoint
for the put me articles article ID public republish operation typically these are written to a http.Request
*/
type PutMeArticlesArticleIDPublicRepublishParams struct {

	/*ArticleID
	  対象記事を指定するために使用

	*/
	ArticleID string
	/*Publish
	  publish object

	*/
	Publish *models.Publish

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put me articles article ID public republish params
func (o *PutMeArticlesArticleIDPublicRepublishParams) WithTimeout(timeout time.Duration) *PutMeArticlesArticleIDPublicRepublishParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put me articles article ID public republish params
func (o *PutMeArticlesArticleIDPublicRepublishParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put me articles article ID public republish params
func (o *PutMeArticlesArticleIDPublicRepublishParams) WithContext(ctx context.Context) *PutMeArticlesArticleIDPublicRepublishParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put me articles article ID public republish params
func (o *PutMeArticlesArticleIDPublicRepublishParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put me articles article ID public republish params
func (o *PutMeArticlesArticleIDPublicRepublishParams) WithHTTPClient(client *http.Client) *PutMeArticlesArticleIDPublicRepublishParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put me articles article ID public republish params
func (o *PutMeArticlesArticleIDPublicRepublishParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArticleID adds the articleID to the put me articles article ID public republish params
func (o *PutMeArticlesArticleIDPublicRepublishParams) WithArticleID(articleID string) *PutMeArticlesArticleIDPublicRepublishParams {
	o.SetArticleID(articleID)
	return o
}

// SetArticleID adds the articleId to the put me articles article ID public republish params
func (o *PutMeArticlesArticleIDPublicRepublishParams) SetArticleID(articleID string) {
	o.ArticleID = articleID
}

// WithPublish adds the publish to the put me articles article ID public republish params
func (o *PutMeArticlesArticleIDPublicRepublishParams) WithPublish(publish *models.Publish) *PutMeArticlesArticleIDPublicRepublishParams {
	o.SetPublish(publish)
	return o
}

// SetPublish adds the publish to the put me articles article ID public republish params
func (o *PutMeArticlesArticleIDPublicRepublishParams) SetPublish(publish *models.Publish) {
	o.Publish = publish
}

// WriteToRequest writes these params to a swagger request
func (o *PutMeArticlesArticleIDPublicRepublishParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param article_id
	if err := r.SetPathParam("article_id", o.ArticleID); err != nil {
		return err
	}

	if o.Publish != nil {
		if err := r.SetBodyParam(o.Publish); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
