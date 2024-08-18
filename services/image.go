package services

import (
	"webback/requests"
	"webback/responses"
)

type Image interface {
	AddImage(request requests.ImagePost, user string) error
	GetImage(id string) (*responses.ImageGet, error)
	GetAllImagesPublic() (*[]responses.ImageGet, error)
}

// Errors

type IncidentalError struct{}

func (e IncidentalError) Error() string {
	return "upload successful but minor error has occurred, certain components may not be available"
}
