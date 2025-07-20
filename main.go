package main

import (
	"auth-rest/cmd"
	"context"
	"log"
	"os"
)

func main() {
	if err := cmd.RootCmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
		return
	}
}
