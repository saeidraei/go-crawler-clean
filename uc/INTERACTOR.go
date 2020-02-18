package uc

import (
	"github.com/saeidraei/go-crawler-clean/domain"
)

// interactor : the struct that will have as properties all the IMPLEMENTED interfaces
// in order to provide them to its methods : the use cases and implement the Handler interface
type interactor struct {
	logger       Logger
	urlValidator UrlValidator
	queueRW      QueueRW
}

// Logger : only used to log stuff
type Logger interface {
	Log(...interface{})
}

type QueueRW interface {
	Enqueue(key string, value domain.Url) error
	Dequeue(key string) (*domain.Url, error)
	All(key string) ([]*domain.Url, error)
}

type UrlValidator interface {
	BeforeCreationCheck(url *domain.Url) error
}
