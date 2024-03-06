package util

import (
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func Ulog(info ...any) {
	log.Println("[DEBUG] ", info)
}

func Ulogf(info ...any) {
	log.Fatalln("[ERROR] ", info)
}

func Uwrite(fileName string, content string) {
	os.WriteFile(fileName, []byte(content), 0644)
}

func MapFind(doc *goquery.Document, rules []string) *goquery.Selection {
	sel := doc.Find(rules[0])
	for i, r := range rules {
		if i == 0 {
			continue
		}
		sel = sel.Find(r)
	}
	return sel
}
