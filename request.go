package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Req2Doc(url string) *goquery.Document {
	Ulog(">>request url:", url)

	res, err := http.Get(url)
	if err != nil {
		Ulogf("request error:", err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		Ulogf("parse html failed:", err)
	}
	return doc
}
