package main

import (
	"github.com/RafaelRMJesus/zlsx-api/db"
	"github.com/RafaelRMJesus/zlsx-api/routes"
)

func main() {
	db.DbConnect()
	routes.InitRoutes()
}
