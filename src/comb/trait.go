package comb

// inner struct
type CombRules struct {
	rules []string
}

type ReqAndParse interface {
	CombineAddress()
	RequestAndParse()
}
