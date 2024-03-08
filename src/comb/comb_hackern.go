package comb

import (
	"query/util"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CombHackern struct {
	CombRules
}

func (comb *CombHackern) CombineAddress() string {
	return comb.address
}

func (comb *CombHackern) RequestAndParse() {
	targetAddress := comb.CombineAddress()

	doc := util.Req2Doc(targetAddress)
	sel := util.MapFind(doc, comb.rules)

	sel.Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if !strings.HasPrefix(href, "from") {
			text := s.Text()
			util.Ulog(text)
		}
	})
}
