package crawlerApi

import (
	"github.com/saeidraei/go-crawler-clean/uc"
	"golang.org/x/net/html"
	"strings"
)
type crawler struct {
}
func New() uc.CrawlerApi {
	return crawler{
	}
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}

func (crawler) GetTitle(htmlString string) (string, bool) {
	var r = strings.NewReader(htmlString)
	doc, err := html.Parse(r)
	if err != nil {
		panic("Fail to parse html")
	}

	return traverse(doc)
}