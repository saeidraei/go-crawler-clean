// +build !netgo

package mock

import (
	"log"

	"github.com/saeidraei/go-crawler-clean/uc"
	"github.com/golang/mock/gomock"
)

// Interactor : is used in order to update its properties accordingly to each test conditions
type Interactor struct {
	Logger           *MockLogger
	UrlRW           *MockUrlRW
	UserValidator    *MockUserValidator
}

type Tester struct {
	Calls      func(*Interactor)
	ShouldPass bool
}

type SimpleLogger struct{}

func (SimpleLogger) Log(logs ...interface{}) {
	log.Println(logs...)
}

//NewMockedInteractor : the Interactor constructor
func NewMockedInteractor(mockCtrl *gomock.Controller) Interactor {
	return Interactor{
		Logger:           NewMockLogger(mockCtrl),
		UserValidator:    NewMockUserValidator(mockCtrl),
		UrlRW:      NewMockUrlRW(mockCtrl),
	}
}

//GetUCHandler : returns a uc.interactor in order to call its methods aka the use cases to test
func (i Interactor) GetUCHandler() uc.Handler {
	return uc.HandlerConstructor{
		Logger:           i.Logger,
		UrlRW:            i.UrlRW,
		UrlValidator:    i.UserValidator,
	}.New()
}
