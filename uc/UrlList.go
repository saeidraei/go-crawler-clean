package uc

import (
	"github.com/saeidraei/go-crawler-clean/domain"
)

func (i interactor) UrlList() ([]*domain.Url, error) {

	urls , err := i.queueRW.All("waiting")
	if err != nil {
		return nil,err
	}
	return urls, nil
}
