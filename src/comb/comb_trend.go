package comb

import (
	"query/util"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	TrendingDaily   string = "daily"
	TrendingWeekly  string = "weekly"
	TrendingMonthly string = "monthly"
)

type CombTrend struct {
	CombRules
	since string
}

func (combTrend *CombTrend) SwitchTrendingType(typ string) {
	combTrend.since = typ
}

func (comb *CombTrend) CombineAddress() string {
	return comb.address + "?since=" + comb.since
}

func (comb *CombTrend) RequestAndParse() {
	targetAddress := comb.CombineAddress()

	doc := util.Req2Doc(targetAddress)
	sel := util.MapFind(doc, comb.rules)

	sel.Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		text = strings.ReplaceAll(text, " ", "")
		text = strings.ReplaceAll(text, "\n", "")
		util.Ulog(text)
	})
}
