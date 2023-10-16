package main

import (
	"github.com/azinudinachzab/hukumonline/app"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app.New().Run()
}
