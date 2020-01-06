package main

import(
  "Web/router"
  "DataBase/connect"
)

func main() {
  db := connect.Data
	defer db.Close()

	var porgs []model.Tests
	db.First(&porgs)
	fmt.Print(porgs)
}
