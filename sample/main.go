package main

import (
	"log"

	"github.com/meeron/go-cms/cms"
)

func main() {
	cmsApp := cms.New()

	log.Fatal(cmsApp.Run(":8080"))
}
