package main

import (
	"fmt"
	"github.com/oguzhanselamoglu/gorest/pkg/api"
	"github.com/oguzhanselamoglu/gorest/pkg/database"
	"github.com/oguzhanselamoglu/gorest/pkg/ui"
)

func main() {
	fmt.Println("...main")

	database.Database()
	api.Api()
	ui.Ui()
}
