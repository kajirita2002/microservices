package main

import (
	"flag"
	"log"
	"microservices/db"

	"github.com/joho/godotenv"
)

var (
	local bool
)

func init() {
	flag.BoolVar(&local, "local", true, "run service local")
	flag.Parse()
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panic(err)
		}
	}

	ctg := db.NewConfig()
	conn, err := db.NewConnection()
}
