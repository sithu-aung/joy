package main

import (
	"log"

	"github.com/sithu-aung/joy/config"
)

func main(){

	log.Println("This is a log message")

	config := config.LoadConfig();

	log.Println(config.DBHost)
}