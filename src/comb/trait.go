package comb

// inner struct
type CombRules struct {
	rules   []string
	address string
}

type CombModal interface {
	CombineAddress() string
	RequestAndParse()
}
