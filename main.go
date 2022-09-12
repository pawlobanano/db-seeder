package main

import (
	"embed"

	"github.com/pawlobanano/db-seeder/seed"
)

//go:embed data
var dataFS embed.FS // Declared here as go:embed doesn't support relative paths (checked 10.09.2022)

func main() {
	seed.DataFS = dataFS
	seed.RunSeeds()
}
