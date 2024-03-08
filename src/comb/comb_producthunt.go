package comb

import (
	"query/util"

	"github.com/PuerkitoBio/goquery"
)

type CombProductHunt struct {
	CombRules
}

func (comb *CombProductHunt) CombineAddress() string {
	return comb.address
}

func (comb *CombProductHunt) RequestAndParse() {
	targetAddress := comb.CombineAddress()

	doc := util.Req2Doc(targetAddress)
	sel := util.MapFind(doc, comb.rules)
	util.Ulog(doc.Text())

	sel.Each(func(i int, s *goquery.Selection) {
		util.Ulog(s.Text())
	})
}
