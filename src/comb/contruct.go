package comb

func NewCombTrend() CombTrend {
	combRules := CombRules{
		rules: []string{"main", ".Box", "article", "h2", "a"},
	}
	return CombTrend{
		since:     TrendingDaily,
		CombRules: combRules,
	}
}

// TODO add filter rule for pass
func NewCombHackern() CombHackern {
	combRules := CombRules{
		rules: []string{"tbody", "tr", ".title", "a"},
	}
	return CombHackern{
		CombRules: combRules,
	}
}
