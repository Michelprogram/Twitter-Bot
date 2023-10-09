package main

import (
	"internal/twitter"
	"internal/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	log.Printf("Start script golang\n")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("No .env file found %s\n", err.Error())
	}

	profile, err := twitter.NewTwitter()

	if err != nil {
		log.Fatalf("Env variable note found %s\n", err.Error())
		os.Exit(1)
	}

	days := utils.DaysBeforeCanada()

	pseudo := utils.GeneratePseudo(days)

	log.Printf("Pseudo generate : %s\n", pseudo)

	err = profile.UpdateTwitterProfile(pseudo)

	if err != nil {
		log.Fatalf("Error while update profile %v\n", err.Error())
		os.Exit(1)
	}

	log.Printf("End script golang \n")
	os.Exit(0)
}
