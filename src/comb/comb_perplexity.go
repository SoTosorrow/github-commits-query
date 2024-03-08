package comb

import (
	"query/util"

	"github.com/PuerkitoBio/goquery"
)

type CombPerplexity struct {
	CombRules
}

func (comb *CombPerplexity) CombineAddress() string {
	return comb.address
}

func (comb *CombPerplexity) RequestAndParse() {
	targetAddress := comb.CombineAddress()

	doc := util.Req2Doc(targetAddress)
	sel := util.MapFind(doc, comb.rules)

	sel.Each(func(i int, s *goquery.Selection) {
		util.Ulog(s.Text())
	})
}
