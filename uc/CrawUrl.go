package uc

import "fmt"

func (i interactor) CrawlUrl(workerId string) {
	url, err := i.queueRW.Dequeue("waiting")
	if err != nil || url == nil {
		crawlerPrint(workerId, "no url to crawl")

		return
	}
	if err != nil {
		panic(err)
	}
	crawlerPrint(workerId, "sending http get request")
	html, err := i.httpRequestApi.GetHtml(url.Address)
	if err != nil {
		crawlerPrint(workerId,
			"could not get response from url:"+url.Address)
		url.FailedCount++
		// if crawl is failed more than 3 times stop crawling
		// (by not adding the url in the waiting queue)
		if url.FailedCount >= 3 {
			err = i.queueRW.Enqueue("failed", url)
			if err != nil {
				panic(err)
			}
		} else {
			//add the failed job to the queue to be executed again
			err = i.queueRW.Enqueue("waiting", url)
			if err != nil {
				panic(err)
			}
		}
		return
	}
	crawlerPrint(workerId, "http get request was successful")
	title, ok := i.crawlerApi.GetTitle(html)
	if ok != true || title == "" {
		crawlerPrint(workerId, "there is no title")
		url.NoTitle = true
	} else {
		crawlerPrint(workerId, "title has been set successfully")
		url.NoTitle = false
		url.Title = title
	}
	err = i.queueRW.Enqueue("done", url)
	if err != nil {
		panic(err)
	}
}

func crawlerPrint(workerId string, message string) {
	fmt.Println("worker" + workerId + " : " + message)
}
