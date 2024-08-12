package services

import (
	"webback/requests"
	"webback/responses"
)

type User interface {
	AddImage(post requests.ImagePost) (*responses.ImagePost, error)
}
