package main

import (
	"fmt"
	"github.com/oguzhanselamoglu/gorest/pkg/database"
	"github.com/oguzhanselamoglu/gorest/pkg/ui"
)

func main() {
	fmt.Println("...main")
	//api.Api()
	database.Database()
	ui.Ui()
}
