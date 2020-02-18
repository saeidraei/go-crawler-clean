package testData

import (
	"github.com/saeidraei/go-crawler-clean/domain"
)


func Url(id string) domain.Url {
	switch id {
	case "aaaaa":
		return a
	default:
		return b
	}
}

var a = domain.Url{
	ID:      "aaaaa",
	Address:     "https://google.com",
}

var b = domain.Url{
	ID:      "bbbbb",
	Address:     "https://saeid.me",
}