package slugger

import (
	"github.com/gosimple/slug"
	"github.com/saeidraei/go-crawler-clean/uc"
)

type slugger struct{}

func New() uc.Slugger {
	return slugger{}
}

func (slugger) NewSlug(initial string) string {
	return slug.Make(initial)
}
