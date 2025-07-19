package main

import (
	"flag"
	"fmt"
	"os"

	"net.craftuniverse.discord/pkg/discord"
)

var APPLICATION_ID int64

func main() {
	flag.Func("help", "Prints this message", func(s string) error {
		flag.PrintDefaults()
		os.Exit(0)

		return nil
	})

	flag.Int64Var(&APPLICATION_ID, "appid", 0, "Discord Client Application ID")
	flag.Parse()

	fmt.Println("Initializing Discord SDK...")

	discord.NewDiscordClient()
}
