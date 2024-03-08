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

// TODO Failed by robot intercept
func NewCombPerplexity() CombPerplexity {
	combRules := CombRules{
		rules:   []string{"main", ".divide-y", ".flex", ".group"},
		address: enum.FullAddressPerplexity,
	}
	return CombPerplexity{
		CombRules: combRules,
	}
}

// TODO Failed by robot intercept
func NewCombProductHunt() CombProductHunt {
	combRules := CombRules{
		rules:   []string{"main", ".homepage-section-0", ".styles_item__Dk_nz", "div"},
		address: enum.FullAddressProductHunt,
	}
	return CombProductHunt{
		CombRules: combRules,
	}
}
