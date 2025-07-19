package main

import (
	"flag"
	"log"

	"net.craftuniverse.discord/pkg/discord"
)

var APPLICATION_ID int64

func main() {
	flag.Int64Var(&APPLICATION_ID, "appid", 0, "Discord Client Application ID")
	flag.Parse()

	log.Println("Initializing Discord SDK...")

	discord.NewDiscordClient()
}
