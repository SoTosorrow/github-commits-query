package main

import (
	"encoding/json"
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
			// TODO but it's still sometimes get error maybe because of the network
			if strings.Contains(s.Text(), "Commits") {
				text = s.Text()
			}
		})
	if !strings.Contains(text, "Commits") {
		Ulogf("parse target failed", text)
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
	// this isn't a good method for find a value from json-string but i dont want to parse the whole json
	// the 13 is becasue the real situation like: currentOid:"asdasdasd...(with 40s id i think)
	currentOidPos := strings.Index(commitsInfo, "currentOid") + 13
	// the 40 is 40 what can i say? maybe it is akind of algorithm but i just count and use it
	// i think it isn't a proper practice to hardcode in functions but i just do it
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

func ParseCommitsBySinglePage(doc *goquery.Document) {
	text := doc.Find(".application-main").Find("main").Find("script").Text()
	// Uwrite("./test", text)
	// this is holy shit to parse the whole json. the structs is fucking too much, i would never do that
	// i think is a good way for query value by loop Index() from the string and maybe it's fast?
	// i just throw it to AI and ask it how to get the "oids" from a unknow format json string. thanks god it works
	// although it is a ugly way and it is in here
	var result map[string]any
	err := json.Unmarshal([]byte(text), &result)
	// Sometimes there will be a {"resolvedServerColorMode":"day"} at the end of the string, what the fuck is this
	if err != nil {
		Ulogf("parse json failed", err)
	}
	oids := make([]string, 0)
	for _, commitGroup := range result["payload"].(map[string]any)["commitGroups"].([]any) {
		for _, commit := range commitGroup.(map[string]any)["commits"].([]any) {
			oids = append(oids, commit.(map[string]any)["oid"].(string))
		}
	}
	// oids = append(oids, result["payload"].(map[string]interface{})["currentCommit"].(map[string]interface{})["oid"].(string))

	Ulog(oids)
}
