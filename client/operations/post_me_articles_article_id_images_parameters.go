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

// NewPostMeArticlesArticleIDImagesParams creates a new PostMeArticlesArticleIDImagesParams object
// with the default values initialized.
func NewPostMeArticlesArticleIDImagesParams() *PostMeArticlesArticleIDImagesParams {
	var ()
	return &PostMeArticlesArticleIDImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeArticlesArticleIDImagesParamsWithTimeout creates a new PostMeArticlesArticleIDImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeArticlesArticleIDImagesParamsWithTimeout(timeout time.Duration) *PostMeArticlesArticleIDImagesParams {
	var ()
	return &PostMeArticlesArticleIDImagesParams{

		timeout: timeout,
	}
}

// NewPostMeArticlesArticleIDImagesParamsWithContext creates a new PostMeArticlesArticleIDImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeArticlesArticleIDImagesParamsWithContext(ctx context.Context) *PostMeArticlesArticleIDImagesParams {
	var ()
	return &PostMeArticlesArticleIDImagesParams{

		Context: ctx,
	}
}

// NewPostMeArticlesArticleIDImagesParamsWithHTTPClient creates a new PostMeArticlesArticleIDImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeArticlesArticleIDImagesParamsWithHTTPClient(client *http.Client) *PostMeArticlesArticleIDImagesParams {
	var ()
	return &PostMeArticlesArticleIDImagesParams{
		HTTPClient: client,
	}
}

/*PostMeArticlesArticleIDImagesParams contains all the parameters to send to the API endpoint
for the post me articles article ID images operation typically these are written to a http.Request
*/
type PostMeArticlesArticleIDImagesParams struct {

	/*ArticleImage
	  article image object

	*/
	ArticleImage *models.ArticleImage
	/*ArticleID
	  対象記事の指定するために使用

	*/
	ArticleID string
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me articles article ID images params
func (o *PostMeArticlesArticleIDImagesParams) WithTimeout(timeout time.Duration) *PostMeArticlesArticleIDImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me articles article ID images params
func (o *PostMeArticlesArticleIDImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me articles article ID images params
func (o *PostMeArticlesArticleIDImagesParams) WithContext(ctx context.Context) *PostMeArticlesArticleIDImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me articles article ID images params
func (o *PostMeArticlesArticleIDImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me articles article ID images params
func (o *PostMeArticlesArticleIDImagesParams) WithHTTPClient(client *http.Client) *PostMeArticlesArticleIDImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me articles article ID images params
func (o *PostMeArticlesArticleIDImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArticleImage adds the articleImage to the post me articles article ID images params
func (o *PostMeArticlesArticleIDImagesParams) WithArticleImage(articleImage *models.ArticleImage) *PostMeArticlesArticleIDImagesParams {
	o.SetArticleImage(articleImage)
	return o
}

// SetArticleImage adds the articleImage to the post me articles article ID images params
func (o *PostMeArticlesArticleIDImagesParams) SetArticleImage(articleImage *models.ArticleImage) {
	o.ArticleImage = articleImage
}

// WithArticleID adds the articleID to the post me articles article ID images params
func (o *PostMeArticlesArticleIDImagesParams) WithArticleID(articleID string) *PostMeArticlesArticleIDImagesParams {
	o.SetArticleID(articleID)
	return o
}

// SetArticleID adds the articleId to the post me articles article ID images params
func (o *PostMeArticlesArticleIDImagesParams) SetArticleID(articleID string) {
	o.ArticleID = articleID
}

// WithContentType adds the contentType to the post me articles article ID images params
func (o *PostMeArticlesArticleIDImagesParams) WithContentType(contentType string) *PostMeArticlesArticleIDImagesParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the post me articles article ID images params
func (o *PostMeArticlesArticleIDImagesParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeArticlesArticleIDImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ArticleImage != nil {
		if err := r.SetBodyParam(o.ArticleImage); err != nil {
			return err
		}
	}

	// path param article_id
	if err := r.SetPathParam("article_id", o.ArticleID); err != nil {
		return err
	}

	// header param content-type
	if err := r.SetHeaderParam("content-type", o.ContentType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
