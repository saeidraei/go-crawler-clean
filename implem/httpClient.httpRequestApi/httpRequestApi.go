package httpRequestApi

import (
	"github.com/saeidraei/go-crawler-clean/uc"
	"io/ioutil"
	"log"
	"net/http"
)
type hra struct {
}
func New() uc.HttpRequestApi {
	return hra{
	}
}


func (hra) GetHtml(url string) (string, error) {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return bodyString, err

	}
	return "",err
}