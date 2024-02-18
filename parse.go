package main

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

/*
1、tbody -> a -> span -> [0] : commit num string like "131 Commits"
*/
func ParseCommitsNum(doc *goquery.Document) int {
	var text string
	doc.Find("tbody").Find("a").Find("span").Each(
		func(i int, s *goquery.Selection) {
			// if text contains "Commits" then is the result
			// why not .First().Text()? sometimes it cant get the corrent result in First one
			if strings.Contains(s.Text(), "Commits") {
				text = s.Text()
			}
		})
	if !strings.Contains(text, "Commits") {
		Ulogf("parse target failed")
	}
	// get 123 from "123 Commits"
	target := strings.Join(
		strings.Split(
			// " Commits" means 8 chars to to discard
			text[:len(text)-8], ","), "")
	commitsNum, err := strconv.Atoi(target)
	if err != nil {
		Ulogf("Commits Num Atoi failed", err)
	}
	return commitsNum
}

/*
#repo-content-pjax-container -> script
*/
func ParseFirstCommitId(doc *goquery.Document) string {
	commitsInfo := doc.Find("#repo-content-pjax-container").Find("script").Text()
	currentOidPos := strings.Index(commitsInfo, "currentOid") + 13
	currentOidNums := 40
	return commitsInfo[currentOidPos : currentOidPos+currentOidNums]
}

/*
react-partial -> div -> button
*/
func ParseBranchName(doc *goquery.Document) string {
	text := doc.Find("react-partial").Find("div").Find("button").First().Text()
	// get "master" from " master"
	return text[2:]
}
