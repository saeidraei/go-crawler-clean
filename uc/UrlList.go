package uc

import (
	"github.com/saeidraei/go-crawler-clean/domain"
)

func (i interactor) UrlList() (map[string][]*domain.Url, error) {

	rU := make(map[string][]*domain.Url)
	urls, err := i.queueRW.All("waiting")
	if err != nil {
		return nil, err
	}
	rU["waiting"] = urls
	urls, err = i.queueRW.All("failed")
	if err != nil {
		return nil, err
	}
	rU["failed"] = urls
	urls, err = i.queueRW.All("done")
	if err != nil {
		return nil, err
	}
	rU["done"] = urls

	return rU, nil
}
