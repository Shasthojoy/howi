package main

import "github.com/okramlabs/howi"

func main() {
	howi := howi.New()

	appMeta := howi.Meta()
	appMeta.SetName("howicli")
	appMeta.SetNamespace("okramlabs")
	appMeta.SetTitle("HOWI CLI")
	appMeta.SetDesc("HOWICLI makes the building of CLI applications in go super fun and easy.")
	appMeta.SetKeywords("golang-tools", "go", "golang-library", "howi")
	appMeta.SetCopyRight(2012, "Marko Kungla")
	appMeta.SetVersion("5.0.0-alpha.1")

	howicli := howi.CLI()
}
