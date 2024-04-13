package main

import (
	"github.com/rs/zerolog"
	"github.com/soloviev1d/avito-banner-service/internal/database"
)

func main() {
	database.PrepareDatabase()
}
