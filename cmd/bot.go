package main

import (
	"internal/twitter"
	"internal/utils"
	"log"
	"os"
)

func main() {

	log.Printf("Start script golang\n")

	profile := twitter.NewTwitter()

	days := utils.DaysBeforeCanada()

	pseudo := utils.GeneratePseudo(days)

	log.Printf("Pseudo generate : %s\n", pseudo)

	err := profile.UpdateTwitterProfile(pseudo)

	if err != nil {
		log.Fatalf("Error while update profile %v\n", err.Error())
		os.Exit(1)
	}

	log.Printf("End script golang \n")
	os.Exit(0)
}
