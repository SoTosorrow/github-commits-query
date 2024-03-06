package comb

import "query/enum"

func NewCombTrend() CombTrend {
	combRules := CombRules{
		rules:   []string{"main", ".Box", "article", "h2", "a"},
		address: enum.FullAddressTrending,
	}
	return CombTrend{
		since:     TrendingDaily,
		CombRules: combRules,
	}
}

// TODO add filter rule for pass
func NewCombHackern() CombHackern {
	combRules := CombRules{
		rules:   []string{"tbody", "tr", ".title", "a"},
		address: enum.FullAddressHackerNews,
	}
	return CombHackern{
		CombRules: combRules,
	}
}
