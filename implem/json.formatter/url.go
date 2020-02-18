package formatter

import (
	"github.com/saeidraei/go-crawler-clean/domain"
)

type Url struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

func NewUrlFromDomain(article domain.Url) Url {

	return Url{
		ID:      article.ID,
		Address: article.Address,
	}
}

func NewUrlsFromDomain(urls ...domain.Url) []Url {
	ret := []Url{} // return at least an empty array (not nil)

	for _, url := range urls {
		ret = append(ret, NewUrlFromDomain(url))
	}

	return ret
}
