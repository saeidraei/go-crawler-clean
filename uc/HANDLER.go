package uc

import (
	"log"

	"github.com/saeidraei/go-crawler-clean/domain"
)

type Handler interface {
	UrlLogic
}

type UrlLogic interface {
	UrlPost(url domain.Url) error
	UrlList() ([]*domain.Url, error)
}

type HandlerConstructor struct {
	Logger       Logger
	UrlValidator UrlValidator
	QueueRW      QueueRW
}

func (c HandlerConstructor) New() Handler {
	if c.Logger == nil {
		log.Fatal("missing Logger")
	}
	if c.QueueRW == nil {
		log.Fatal("missing QueueRW")
	}

	return interactor{
		logger:       c.Logger,
		queueRW:      c.QueueRW,
		urlValidator: c.UrlValidator,
	}
}
