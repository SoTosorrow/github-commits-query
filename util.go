package main

import (
	"log"
	"os"
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
