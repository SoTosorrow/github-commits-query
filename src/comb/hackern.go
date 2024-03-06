package comb

import (
	"query/util"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CombHackern struct {
	CombRules
}

func (combHackern *CombHackern) CombineAddress() string {
	return combHackern.address
}

func (combHackern *CombHackern) RequestAndParse() {
	targetAddress := combHackern.CombineAddress()

	doc := util.Req2Doc(targetAddress)
	sel := util.MapFind(doc, combHackern.rules)

	sel.Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if !strings.HasPrefix(href, "from") {
			text := s.Text()
			util.Ulog(text)
		}
	})
}
