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
