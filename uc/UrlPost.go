package uc

import (
	"github.com/saeidraei/go-crawler-clean/domain"
	"math/rand"
	"strings"
	"time"
)

func (i interactor) UrlPost(url domain.Url) error {

	//if there is a validator validate the url
	if i.urlValidator != nil {
		if err := i.urlValidator.BeforeCreationCheck(&url); err != nil {
			return err
		}
	}

	url.ID = randStringBytes(7)
	//add the http if it didn't exist
	if !strings.Contains(url.Address,"://"){
		url.Address = "http://" + url.Address
	}
	err := i.queueRW.Enqueue("waiting",&url)
	if err != nil {
		return err
	}

	return  nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
