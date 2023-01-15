package main

import (
	"github.com/lu-css/go-url-encounter/pkg/database"
	"github.com/lu-css/go-url-encounter/pkg/server"
)

func main() {
  database.Connect()
  server.Start()
}
