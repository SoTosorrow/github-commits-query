package comb

import (
	"query/enum"
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

func (combTrend *CombTrend) CombineAddress() string {
	return enum.FullAddressTrending + "?since=" + combTrend.since
}

func (combTrend *CombTrend) RequestAndParse() {
	targetAddress := combTrend.CombineAddress()

	doc := util.Req2Doc(targetAddress)
	sel := util.MapFind(doc, combTrend.rules)

	sel.Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		text = strings.ReplaceAll(text, " ", "")
		text = strings.ReplaceAll(text, "\n", "")
		util.Ulog(text)
	})
}
