// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/alisista/go-alis-client/models"
)

// GetArticlesPopularReader is a Reader for the GetArticlesPopular structure.
type GetArticlesPopularReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArticlesPopularReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetArticlesPopularOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetArticlesPopularOK creates a GetArticlesPopularOK with default headers values
func NewGetArticlesPopularOK() *GetArticlesPopularOK {
	return &GetArticlesPopularOK{}
}

/*GetArticlesPopularOK handles this case with default header values.

人気記事一覧
*/
type GetArticlesPopularOK struct {
	Payload []*models.ArticleInfo
}

func (o *GetArticlesPopularOK) Error() string {
	return fmt.Sprintf("[GET /articles/popular][%d] getArticlesPopularOK  %+v", 200, o.Payload)
}

func (o *GetArticlesPopularOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
