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
	UrlList() (map[string][]*domain.Url, error)
	CrawlUrl(workerId string)
}

type HandlerConstructor struct {
	Logger       Logger
	UrlValidator UrlValidator
	QueueRW      QueueRW
	CrawlerApi   CrawlerApi
	HttpRequestApi   HttpRequestApi
}

func (c HandlerConstructor) New() Handler {
	if c.Logger == nil {
		log.Fatal("missing Logger")
	}
	if c.QueueRW == nil {
		log.Fatal("missing QueueRW")
	}
	if c.CrawlerApi == nil {
		log.Fatal("missing CrawlerApi")
	}

	return interactor{
		logger:       c.Logger,
		queueRW:      c.QueueRW,
		urlValidator: c.UrlValidator,
		crawlerApi:   c.CrawlerApi,
		httpRequestApi: c.HttpRequestApi,
	}
}
