package main

import (
	"log"

	"github.com/meeron/go-cms/cms"
)

func main() {
	cmsApp := cms.New(&cms.CmsConfig{})

	log.Fatal(cmsApp.Run(":8080"))
}
