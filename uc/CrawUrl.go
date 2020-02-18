package uc

func (i interactor) CrawlUrl() (string, error) {
	url, err := i.queueRW.Dequeue("waiting")
	if err != nil || url == nil{
		return "", err
	}
	html, err := i.httpRequestApi.GetHtml(url.Address)
	title, ok := i.crawlerApi.GetTitle(html)
	if ok != true {
		return "", err
	}
	return title, nil
}
