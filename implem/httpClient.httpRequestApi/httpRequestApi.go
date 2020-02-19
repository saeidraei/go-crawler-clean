package httpRequestApi

import (
	"github.com/saeidraei/go-crawler-clean/uc"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"time"
)

type hra struct {
}

func New() uc.HttpRequestApi {
	return hra{
	}
}

func (hra) GetHtml(url string) (string, error) {
	var client http.Client
	timeout := viper.GetInt("crawler.requestTimeout")
	client.Timeout = time.Duration(timeout) * time.Second
	resp, err := client.Get(url)
	if err != nil {
		return "", err
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
	return "", err
}
