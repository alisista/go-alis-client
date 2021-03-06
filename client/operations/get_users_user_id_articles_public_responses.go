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

// GetUsersUserIDArticlesPublicReader is a Reader for the GetUsersUserIDArticlesPublic structure.
type GetUsersUserIDArticlesPublicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUserIDArticlesPublicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsersUserIDArticlesPublicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsersUserIDArticlesPublicOK creates a GetUsersUserIDArticlesPublicOK with default headers values
func NewGetUsersUserIDArticlesPublicOK() *GetUsersUserIDArticlesPublicOK {
	return &GetUsersUserIDArticlesPublicOK{}
}

/*GetUsersUserIDArticlesPublicOK handles this case with default header values.

公開記事一覧
*/
type GetUsersUserIDArticlesPublicOK struct {
	Payload []*models.ArticleInfo
}

func (o *GetUsersUserIDArticlesPublicOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/articles/public][%d] getUsersUserIdArticlesPublicOK  %+v", 200, o.Payload)
}

func (o *GetUsersUserIDArticlesPublicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
